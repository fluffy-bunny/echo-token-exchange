package token

import (
	"context"
	contracts_config "echo-starter/internal/contracts/config"
	contracts_stores_apiresources "echo-starter/internal/contracts/stores/apiresources"
	contracts_clients "echo-starter/internal/contracts/stores/clients"
	contracts_stores_keymaterial "echo-starter/internal/contracts/stores/keymaterial"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	models "echo-starter/internal/models"
	echo_oauth2 "echo-starter/internal/services/go-oauth2/oauth2"
	utils "echo-starter/internal/utils"
	wellknown "echo-starter/internal/wellknown"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	di "github.com/dozm/di"
	"github.com/fatih/structs"
	fluffycore_contracts_common "github.com/fluffy-bunny/fluffycore/contracts/common"
	contracts_handler "github.com/fluffy-bunny/fluffycore/echo/contracts/handler"
	core_hashset "github.com/fluffy-bunny/fluffycore/gods/sets/hashset"
	core_utils "github.com/fluffy-bunny/fluffycore/utils"
	oauth2 "github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	oauth2_models "github.com/go-oauth2/oauth2/v4/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

type (
	service struct {
		Now                  fluffycore_contracts_common.TimeNow           `inject:""`
		Config               *contracts_config.Config                      `inject:""`
		APIResources         contracts_stores_apiresources.IAPIResources   `inject:""`
		KeyMaterial          contracts_stores_keymaterial.IKeyMaterial     `inject:""`
		JwtTokenStore        contracts_stores_tokenstore.IJwtTokenStore    `inject:""`
		ClientRequest        contracts_clients.IClientRequest              `inject:""`
		TokenHandlerAccessor contracts_tokenhandlers.ITokenHandlerAccessor `inject:""`
		ReferenceTokenStore  contracts_stores_tokenstore.ITokenStore       `inject:""`
		TokenHandler         contracts_tokenhandlers.ITokenHandler
		accessGenerate       echo_oauth2.AccessGenerate
		signingKey           *models.SigningKey
		issuer               string
	}
)

var stemService *service = new(service)

func init() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

func (s *service) Ctor(
	now fluffycore_contracts_common.TimeNow,
	config *contracts_config.Config,
	apiResources contracts_stores_apiresources.IAPIResources,
	keyMaterial contracts_stores_keymaterial.IKeyMaterial,
	jwtTokenStore contracts_stores_tokenstore.IJwtTokenStore,
	clientRequest contracts_clients.IClientRequest,
	tokenHandlerAccessor contracts_tokenhandlers.ITokenHandlerAccessor,
	referenceTokenStore contracts_stores_tokenstore.ITokenStore,
) (*service, error) {
	obj := &service{
		Now:                  now,
		Config:               config,
		APIResources:         apiResources,
		KeyMaterial:          keyMaterial,
		JwtTokenStore:        jwtTokenStore,
		ClientRequest:        clientRequest,
		TokenHandlerAccessor: tokenHandlerAccessor,
		ReferenceTokenStore:  referenceTokenStore,
	}
	obj.TokenHandler = obj.TokenHandlerAccessor.GetTokenHandler()
	signingKey, err := obj.KeyMaterial.GetSigningKey()
	if err != nil {
		panic(err)
	}
	obj.signingKey = signingKey
	return obj, nil
}

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder di.ContainerBuilder) {
	contracts_handler.AddScopedIHandleWithMetadata[*service](builder,
		stemService.Ctor,
		[]contracts_handler.HTTPVERB{
			contracts_handler.POST,
		},
		wellknown.OAuth2TokenPath,
	)

}

func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (s *service) Do(c echo.Context) error {
	rootPath := utils.GetMyRootPath(c)
	s.issuer = rootPath + "/"

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

	claims.Set("client_id", client.ClientID)

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
		data[models.TokenTypeRefreshToken] = refresh
	}

	return data
}

// GenerateAccessToken generate the access token
func (s *service) GenerateAccessToken(ctx context.Context,
	validatedResult *contracts_tokenhandlers.ValidatedTokenRequestResult,
	subject string,
	claims models.IClaims) (oauth2.TokenInfo, error) {
	now := s.Now()
	//------------------------------------------------------------------------
	client := s.ClientRequest.GetClient()

	ti := oauth2_models.NewToken()
	ti.SetClientID(client.ClientID)
	ti.SetUserID(subject)

	scope := validatedResult.Params["scope"]
	ti.SetScope(scope)
	scopes := strings.Split(scope, " ")
	scopeSet := core_hashset.NewStringSet(scopes...)
	createAt := now
	ti.SetAccessCreateAt(createAt)

	ti.SetAccessExpiresIn(time.Duration(client.AccessTokenLifetime) * time.Second)

	expiresAt := createAt.Add(time.Second * time.Duration(client.AccessTokenLifetime))
	standardClaims := &jwt.StandardClaims{
		IssuedAt:  createAt.Unix(),
		ExpiresAt: expiresAt.Unix(),
		Issuer:    s.issuer,
		Audience:  client.ClientID,
		Subject:   subject,
		Id:        xid.New().String(),
	}
	var buildClaimsMap = func(ctx context.Context, standardClaims *jwt.StandardClaims, extras models.IClaims) models.IClaims {
		audienceSet := core_hashset.NewStringSet()
		if !core_utils.IsEmptyOrNil(standardClaims.Audience) {
			audienceSet.Add(standardClaims.Audience)
		}
		if !core_utils.IsNil(extras) {
			extraAudInterface := extras.Get("aud")
			switch extraAudienceType := extraAudInterface.(type) {
			case string:
				audienceSet.Add(extraAudienceType)
			case []string:
				audienceSet.Add(extraAudienceType...)
			}
		}
		extras.Set("aud", audienceSet.Values())

		var standard map[string]interface{}
		standardJSON, _ := json.Marshal(standardClaims)
		json.Unmarshal(standardJSON, &standard)
		delete(standard, "aud")

		for k, v := range standard {
			extras.Set(k, v)
		}
		return extras
	}
	claims = buildClaimsMap(ctx, standardClaims, claims)

	var err error
	var tokenHandle string
	if client.AccessTokenType == models.Reference {
		handle := utils.GenerateHandle()
		tokenHandle, err = s.ReferenceTokenStore.StoreToken(ctx, handle, &models.TokenInfo{
			Metadata: models.TokenMetadata{
				Type:        models.TokenTypeReferenceToken,
				ClientID:    client.ClientID,
				Subject:     subject,
				Expiration:  expiresAt,
				IssedAt:     now,
				Issuer:      s.issuer,
				OrgID:       "TODO", // TODO pull this from the claims as it is integral to the subject
				IntegrityID: "TODO", // TODO pull this from the claims as it is integral to the subject
			},
			Data: claims.Claims(),
		})
		if err != nil {
			return nil, err
		}

	} else {
		tokenHandle, err = s.JwtTokenStore.MintToken(ctx, claims)
		if err != nil {
			return nil, err
		}
	}

	ti.SetAccess(tokenHandle)

	if client.AllowOfflineAccess && scopeSet.Contains("offline_access") {
		if core_utils.IsEmptyOrNil(validatedResult.RefreshTokenHandle) {
			panic("refresh token handle is empty") // fix your code
		}
		var absoluteExpiration = now.Add(time.Second * time.Duration(client.AbsoluteRefreshTokenLifetime))
		if client.AbsoluteRefreshTokenLifetime <= 0 {
			absoluteExpiration = now.Add(time.Hour * 24 * 365 * 10) // 10 years
		}
		var expiration = now.Add(time.Second * time.Duration(client.RefreshTokenExpiration))
		rtInfo := &models.RefreshTokenInfo{
			ClientID:                     client.ClientID,
			Subject:                      subject,
			Scope:                        scope,
			GrantType:                    validatedResult.GrantType,
			Expiration:                   expiration,
			AbsoluteExpiration:           absoluteExpiration,
			Params:                       validatedResult.Params,
			RefreshTokenGraceEnabled:     client.RefreshTokenGraceEnabled,
			RefreshTokenGraceTTL:         client.RefreshTokenGraceTTL,
			RefreshTokenGraceMaxAttempts: client.RefreshTokenGraceMaxAttempts,
			RefreshTokenGraceAttempts:    0,
		}

		data := structs.Map(rtInfo)
		handle, err := s.ReferenceTokenStore.StoreToken(ctx,
			validatedResult.RefreshTokenHandle,
			&models.TokenInfo{
				Metadata: models.TokenMetadata{
					Type:       models.TokenTypeRefreshToken,
					ClientID:   client.ClientID,
					Subject:    subject,
					Expiration: expiration,
					IssedAt:    now,
					Issuer:     s.issuer,
				},
				Data: data,
			})

		if err != nil {
			return nil, err
		}
		ti.SetRefresh(handle)
	}

	return ti, nil
}
