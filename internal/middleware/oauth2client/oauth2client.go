package claimsprincipal

import (
	"echo-starter/internal/models"
	"strings"

	contracts_clients "echo-starter/internal/contracts/clients"
	"echo-starter/internal/utils"
	"echo-starter/internal/wellknown"

	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"

	core_wellknown "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/wellknown"
	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-oauth2/oauth2/v4/errors"
	oauth2_server "github.com/go-oauth2/oauth2/v4/server"
	"github.com/labstack/echo/v4"
)

func _clientInfoHandler(c echo.Context, clientStore contracts_clients.IClientStore) (client *models.Client, err error) {
	r := c.Request()
	clientID, clientSecret, err := oauth2_server.ClientBasicHandler(r)
	if err != nil {
		clientID, clientSecret, err = oauth2_server.ClientFormHandler(r)
	}
	if err != nil {
		return nil, err
	}

	client, _, _ = clientStore.GetClient(r.Context(), clientID)
	var match bool
	for _, sc := range client.ClientSecrets {
		match, _ = utils.ComparePasswordHash(clientSecret, sc.Value)
		if match {
			break
		}
	}
	if !match {
		err = errors.ErrInvalidClient
	}
	if err != nil {
		client = nil
	}
	return
}

func AuthenticateOAuth2Client(root di.Container) echo.MiddlewareFunc {

	clientStore := contracts_clients.GetIClientStoreFromContainer(root)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			var grantType string
			path := r.URL.Path
			switch path {
			case wellknown.OAuth2TokenPath:
				grantType = r.FormValue("grant_type")
				if !wellknown.SupportedGrantTypes.Contains(grantType) {
					return c.JSON(401, "grant_type not supported")
				}
			case wellknown.OAuth2RevokePath:
			default:
				return next(c)
			}
			client, err := _clientInfoHandler(c, clientStore)
			if err != nil {
				return c.JSON(401, "client_id or client_secret is invalid")
			}
			switch c.Request().URL.Path {
			case wellknown.OAuth2TokenPath:
				scope := r.FormValue("scope")
				if !core_utils.IsEmptyOrNil(scope) {
					requestedScopes := strings.Split(scope, " ")
					requestedScopeSet := core_hashset.NewStringSet(requestedScopes...)
					if !client.AllowedScopesSet.Contains(requestedScopeSet.Values()...) {
						return c.JSON(401, "scope is invalid")
					}
				}
				ok := client.AllowedGrantTypesSet.Contains(grantType)
				if !ok {
					return c.JSON(401, "grant_type is invalid")
				}

			}

			// validate that the required form arguments are present
			switch grantType {
			case wellknown.OAuth2GrantType_RefreshToken:
				refreshToken := r.FormValue("refresh_token")
				if core_utils.IsEmptyOrNil(refreshToken) {
					return c.JSON(401, "refresh_token is required")
				}
			case wellknown.OAuth2GrantType_TokenExchange:
				subjectToken := r.FormValue("subject_token")
				if core_utils.IsEmptyOrNil(subjectToken) {
					return c.JSON(401, "subject_token is required")
				}
				subjectTokenType := r.FormValue("subject_token_type")
				if core_utils.IsEmptyOrNil(subjectTokenType) {
					return c.JSON(401, "subject_token_type is required")
				}
			}

			scopedContainer := c.Get(core_wellknown.SCOPED_CONTAINER_KEY).(di.Container)
			clientRequest := contracts_clients.GetIClientRequestInternalFromContainer(scopedContainer)
			clientRequest.SetClient(client)

			tokenHandlerAccessor := contracts_tokenhandlers.GetIInternalTokenHandlerAccessorFromContainer(scopedContainer)
			tokenHandlerAccessor.SetGrantType(grantType)
			switch grantType {
			case wellknown.OAuth2GrantType_ClientCredentials:
				tokenHandler := contracts_tokenhandlers.GetIClientCredentialsTokenHandlerFromContainer(scopedContainer)
				tokenHandlerAccessor.SetTokenHandler(tokenHandler)
			case wellknown.OAuth2GrantType_RefreshToken:
				tokenHandler := contracts_tokenhandlers.GetIRefreshTokenHandlerFromContainer(scopedContainer)
				tokenHandlerAccessor.SetTokenHandler(tokenHandler)
			case wellknown.OAuth2GrantType_TokenExchange:
				tokenHandler := contracts_tokenhandlers.GetITokenExchangeTokenHandlerFromContainer(scopedContainer)
				tokenHandlerAccessor.SetTokenHandler(tokenHandler)
			}
			return next(c)
		}
	}
}
