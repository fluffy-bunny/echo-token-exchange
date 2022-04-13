package token

import (
	"context"
	contracts_config "echo-starter/internal/contracts/config"
	contracts_stores_apiresources "echo-starter/internal/contracts/stores/apiresources"
	contracts_clients "echo-starter/internal/contracts/stores/clients"
	contracts_stores_jwttoken "echo-starter/internal/contracts/stores/jwttoken"
	contracts_stores_keymaterial "echo-starter/internal/contracts/stores/keymaterial"

	contracts_stores_referencetoken "echo-starter/internal/contracts/stores/referencetoken"
	contracts_stores_refreshtoken "echo-starter/internal/contracts/stores/refreshtoken"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/models"
	echo_oauth2 "echo-starter/internal/services/go-oauth2/oauth2"
	"echo-starter/internal/utils"
	"echo-starter/internal/wellknown"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	oauth2_models "github.com/go-oauth2/oauth2/v4/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

type (
	service struct {
		Config               *contracts_config.Config                             `inject:""`
		Logger               contracts_logger.ILogger                             `inject:""`
		ClientStore          contracts_clients.IClientStore                       `inject:""`
		APIResources         contracts_stores_apiresources.IAPIResources          `inject:""`
		KeyMaterial          contracts_stores_keymaterial.IKeyMaterial            `inject:""`
		JwtTokenStore        contracts_stores_jwttoken.IJwtTokenStore             `inject:""`
		ClientRequest        contracts_clients.IClientRequest                     `inject:""`
		TokenHandlerAccessor contracts_tokenhandlers.ITokenHandlerAccessor        `inject:""`
		RefreshTokenStore    contracts_stores_refreshtoken.IRefreshTokenStore     `inject:""`
		ReferenceTokenStore  contracts_stores_referencetoken.IReferenceTokenStore `inject:""`
		TokenHandler         contracts_tokenhandlers.ITokenHandler
		accessGenerate       echo_oauth2.AccessGenerate
		signingKey           *models.SigningKey
		issuer               string
	}
)

func assertImplementation() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder *di.Builder) {
	contracts_handler.AddScopedIHandlerEx(builder,
		reflectType,
		[]contracts_handler.HTTPVERB{
			contracts_handler.POST,
		},
		wellknown.OAuth2TokenPath)
}
func (s *service) Ctor() {
	s.TokenHandler = s.TokenHandlerAccessor.GetTokenHandler()

	signingKey, err := s.KeyMaterial.GetSigningKey()
	if err != nil {
		panic(err)
	}
	s.signingKey = signingKey

}
func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (s *service) Do(c echo.Context) error {
	rootPath := utils.GetMyRootPath(c)
	s.issuer = rootPath

	return s.processRequest(c)
}
func getMyRootPath(c echo.Context) string {
	return fmt.Sprintf("%s://%s", c.Scheme(), c.Request().Host)
}

func (s *service) processRequest(c echo.Context) error {
	ctx := c.Request().Context()
	r := c.Request()
	w := c.Response()
	client := s.ClientRequest.GetClient()
	validatedResult, err := s.TokenHandler.ValidationTokenRequest(r)
	if err != nil {
		return s.tokenError(c.Response(), err)
	}
	validatedResult.ClientID = client.ClientID
	claims, err := s.TokenHandler.ProcessTokenRequest(ctx, validatedResult)
	if err != nil {
		return s.tokenError(c.Response(), err)
	}
	iClaims := claims.(models.IClaims)

	iClaims.Set("client_id", client.ClientID)

	ti, err := s.GetAccessToken(ctx, validatedResult, "", claims)
	if err != nil {
		return s.tokenError(w, err)
	}

	return s.token(w, s.GetTokenData(ti), nil)

}

func (s *service) tokenError(w http.ResponseWriter, err error) error {
	data, statusCode, header := s.GetErrorData(err)
	return s.token(w, data, header, statusCode)
}
func (s *service) token(w http.ResponseWriter, data map[string]interface{}, header http.Header, statusCode ...int) error {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")

	for key := range header {
		w.Header().Set(key, header.Get(key))
	}

	status := http.StatusOK
	if len(statusCode) > 0 && statusCode[0] > 0 {
		status = statusCode[0]
	}

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
func (s *service) GetErrorData(err error) (map[string]interface{}, int, http.Header) {
	var re errors.Response
	if v, ok := errors.Descriptions[err]; ok {
		re.Error = err
		re.Description = v
		re.StatusCode = errors.StatusCodes[err]
	} else {

		if re.Error == nil {
			re.Error = errors.ErrServerError
			re.Description = errors.Descriptions[errors.ErrServerError]
			re.StatusCode = errors.StatusCodes[errors.ErrServerError]
		}
	}

	data := make(map[string]interface{})
	if err := re.Error; err != nil {
		data["error"] = err.Error()
	}

	if v := re.ErrorCode; v != 0 {
		data["error_code"] = v
	}

	if v := re.Description; v != "" {
		data["error_description"] = v
	}

	if v := re.URI; v != "" {
		data["error_uri"] = v
	}

	statusCode := http.StatusInternalServerError
	if v := re.StatusCode; v > 0 {
		statusCode = v
	}

	return data, statusCode, re.Header
}

// GetAccessToken access token
func (s *service) GetAccessToken(ctx context.Context,
	validatedResult *contracts_tokenhandlers.ValidatedTokenRequestResult,
	subject string, claims models.IClaims) (oauth2.TokenInfo,
	error) {

	switch validatedResult.GrantType {
	case wellknown.OAuth2GrantType_ClientCredentials:
		return s.GenerateAccessToken(ctx, validatedResult, subject, claims)
	case wellknown.OAuth2GrantType_RefreshToken:
		return s.GenerateAccessToken(ctx, validatedResult, subject, claims)
	}

	return nil, errors.ErrUnsupportedGrantType
}

// GetTokenData token data
func (s *service) GetTokenData(ti oauth2.TokenInfo) map[string]interface{} {
	data := map[string]interface{}{
		"access_token": ti.GetAccess(),
		"token_type":   s.Config.TokenType,
		"expires_in":   int64(ti.GetAccessExpiresIn() / time.Second),
	}

	if scope := ti.GetScope(); scope != "" {
		data["scope"] = scope
	}

	if refresh := ti.GetRefresh(); refresh != "" {
		data["refresh_token"] = refresh
	}

	return data
}

// GenerateAccessToken generate the access token
func (s *service) GenerateAccessToken(ctx context.Context,
	validatedResult *contracts_tokenhandlers.ValidatedTokenRequestResult,
	subject string,
	claims models.IClaims) (oauth2.TokenInfo, error) {

	//------------------------------------------------------------------------
	client := s.ClientRequest.GetClient()

	ti := oauth2_models.NewToken()
	ti.SetClientID(client.ClientID)
	ti.SetUserID(subject)

	scope, _ := validatedResult.Params["scope"]
	ti.SetScope(scope)
	scopes := strings.Split(scope, " ")
	scopeSet := core_hashset.NewStringSet(scopes...)
	createAt := time.Now()
	ti.SetAccessCreateAt(createAt)

	ti.SetAccessExpiresIn(time.Duration(client.AccessTokenLifetime) * time.Second)

	standardClaims := &jwt.StandardClaims{
		IssuedAt:  createAt.Unix(),
		ExpiresAt: createAt.Add(time.Second * time.Duration(client.AccessTokenLifetime)).Unix(),
		Issuer:    s.issuer,
		Audience:  client.ClientID,
		Subject:   subject,
		Id:        xid.New().String(),
	}
	jwtToken, err := s.JwtTokenStore.MintToken(ctx, standardClaims, claims)
	if err != nil {
		return nil, err
	}

	ti.SetAccess(jwtToken)

	if client.AllowOfflineAccess && scopeSet.Contains("offline_access") {
		var absoluteExpiration = time.Now().Add(time.Second * time.Duration(client.AbsoluteRefreshTokenLifetime))
		if client.AbsoluteRefreshTokenLifetime <= 0 {
			absoluteExpiration = time.Now().Add(time.Hour * 24 * 365 * 10) // 10 years
		}
		var expiration = time.Now().Add(time.Second * time.Duration(client.RefreshTokenExpiration))
		handle, err := s.RefreshTokenStore.StoreRefreshToken(ctx,
			&contracts_stores_refreshtoken.RefreshTokenInfo{
				ClientID:           client.ClientID,
				Subject:            subject,
				Scope:              scope,
				GrantType:          validatedResult.GrantType,
				Expiration:         expiration,
				AbsoluteExpiration: absoluteExpiration,
				Params:             validatedResult.Params,
			})
		if err != nil {
			return nil, err
		}
		ti.SetRefresh(handle)
	}

	return ti, nil
}
