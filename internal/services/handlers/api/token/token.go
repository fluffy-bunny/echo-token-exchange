package token

import (
	"context"
	contracts_apiresources "echo-starter/internal/contracts/apiresources"
	contracts_clients "echo-starter/internal/contracts/clients"
	contracts_config "echo-starter/internal/contracts/config"
	contracts_go_oauth2_oauth2 "echo-starter/internal/contracts/go-oauth2/oauth2"
	"echo-starter/internal/models"
	"echo-starter/internal/services/go-oauth2/oauth2/generates"
	"echo-starter/internal/utils"
	"echo-starter/internal/wellknown"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	contracts_stores_refreshtoken "echo-starter/internal/contracts/stores/refreshtoken"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	echo_oauth2 "echo-starter/internal/services/go-oauth2/oauth2"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	oauth2_models "github.com/go-oauth2/oauth2/v4/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type (
	service struct {
		Config               *contracts_config.Config                         `inject:""`
		Logger               contracts_logger.ILogger                         `inject:""`
		ClientStore          contracts_clients.IClientStore                   `inject:""`
		APIResources         contracts_apiresources.IAPIResources             `inject:""`
		SigningKeyStore      contracts_go_oauth2_oauth2.ISigningKeyStore      `inject:""`
		ClientRequest        contracts_clients.IClientRequest                 `inject:""`
		TokenHandlerAccessor contracts_tokenhandlers.ITokenHandlerAccessor    `inject:""`
		RefreshTokenStore    contracts_stores_refreshtoken.IRefreshTokenStore `inject:""`
		TokenHandler         contracts_tokenhandlers.ITokenHandler
		accessGenerate       echo_oauth2.AccessGenerate
		signingKey           *models.SigningKey
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

	signingKey, err := s.SigningKeyStore.GetSigningKey()
	if err != nil {
		panic(err)
	}
	s.signingKey = signingKey
	/*
		privateKey, publicKey, err := ecdsa.DecodePrivatePem(signingKey.Password, signingKey.PrivateKey)
		if err != nil {
			panic(err)
		}
		encPriv, _, err := ecdsa.Encode("", privateKey, publicKey)
	*/

}
func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (s *service) Do(c echo.Context) error {
	rootPath := utils.GetMyRootPath(c)
	jwtGenerator := generates.NewJWTAccessGenerate(s.signingKey.Kid, []byte(s.signingKey.PrivateKey), jwt.SigningMethodES256)
	jwtGenerator.Issuer = rootPath
	s.accessGenerate = jwtGenerator
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
	fmt.Println("claims", claims)
	_, tgr, err := s.ValidationTokenRequest(c.Request())

	ti, err := s.GetAccessToken(ctx, validatedResult, tgr, claims)
	if err != nil {
		return s.tokenError(w, err)
	}

	return s.token(w, s.GetTokenData(ti), nil)

}
func (s *service) ValidationTokenRequest(r *http.Request) (oauth2.GrantType, *oauth2.TokenGenerateRequest, error) {
	// grant_type and scopes have been validated in the middleware
	gt := oauth2.GrantType(r.FormValue("grant_type"))
	client := s.ClientRequest.GetClient()

	tgr := &oauth2.TokenGenerateRequest{
		ClientID: client.ClientID,
		Request:  r,
	}

	switch gt {
	case oauth2.ClientCredentials:
		tgr.Scope = r.FormValue("scope")
	case oauth2.Refreshing:
		tgr.Refresh = r.FormValue("refresh_token")
		tgr.Scope = r.FormValue("scope")
		if tgr.Refresh == "" {
			return "", nil, errors.ErrInvalidRequest
		}
	case "urn:ietf:params:oauth:grant-type:token-exchange":
	default:
		return "", nil, errors.ErrUnsupportedGrantType
	}

	return gt, tgr, nil
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

// CheckGrantType check allows grant type
func (s *service) CheckGrantType(gt oauth2.GrantType) bool {
	for _, agt := range s.Config.AllowedGrantTypes {
		if agt == gt {
			return true
		}
	}
	return false
}

// GetAccessToken access token
func (s *service) GetAccessToken(ctx context.Context,
	validatedResult *contracts_tokenhandlers.ValidatedTokenRequestResult,
	tgr *oauth2.TokenGenerateRequest, claims contracts_tokenhandlers.Claims) (oauth2.TokenInfo,
	error) {

	switch validatedResult.GrantType {
	case wellknown.OAuth2GrantType_ClientCredentials:
		return s.GenerateAccessToken(ctx, validatedResult, tgr, claims)
	case wellknown.OAuth2GrantType_RefreshToken:
		return s.GenerateAccessToken(ctx, validatedResult, tgr, claims)
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
	tgr *oauth2.TokenGenerateRequest,
	claims contracts_tokenhandlers.Claims) (oauth2.TokenInfo, error) {

	//------------------------------------------------------------------------
	client := s.ClientRequest.GetClient()

	ti := oauth2_models.NewToken()
	ti.SetClientID(tgr.ClientID)
	ti.SetUserID(tgr.UserID)
	ti.SetRedirectURI(tgr.RedirectURI)
	scope, _ := validatedResult.Params["scope"]
	ti.SetScope(scope)
	scopes := strings.Split(scope, " ")
	scopeSet := core_hashset.NewStringSet(scopes...)
	createAt := time.Now()
	ti.SetAccessCreateAt(createAt)

	ti.SetAccessExpiresIn(time.Duration(client.AccessTokenLifetime) * time.Second)

	td := &echo_oauth2.GenerateBasic{
		APIResources: s.APIResources,
		Client:       client,
		UserID:       tgr.UserID,
		CreateAt:     createAt,
		TokenInfo:    ti,
		Request:      tgr.Request,
	}

	av, err := s.accessGenerate.Token(ctx, td, claims)
	if err != nil {
		return nil, err
	}
	ti.SetAccess(av)

	if client.AllowOfflineAccess && scopeSet.Contains("offline_access") {
		var absoluteExpiration = time.Now().Add(time.Second * time.Duration(client.AbsoluteRefreshTokenLifetime))
		if client.AbsoluteRefreshTokenLifetime <= 0 {
			absoluteExpiration = time.Now().Add(time.Hour * 24 * 365 * 10) // 10 years
		}
		var expiration = time.Now().Add(time.Second * time.Duration(client.RefreshTokenExpiration))
		handle, err := s.RefreshTokenStore.StoreRefreshToken(ctx,
			&contracts_stores_refreshtoken.RefreshTokenInfo{
				ClientID:           tgr.ClientID,
				Subject:            tgr.UserID,
				Scope:              tgr.Scope,
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
