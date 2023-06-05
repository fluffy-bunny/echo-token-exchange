package discovery

import (
	"echo-starter/internal/utils"
	"echo-starter/internal/wellknown"
	"net/http"
	"reflect"

	"echo-starter/internal/models"

	di "github.com/dozm/di"
	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	"github.com/labstack/echo/v4"
)

type (
	service struct{}
)

func assertImplementation() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder di.ContainerBuilder) {
	contracts_handler.AddScopedIHandlerEx(builder,
		reflectType,
		[]contracts_handler.HTTPVERB{
			contracts_handler.GET,
		},
		wellknown.WellKnownOpenIDCOnfiguationPath)
}

func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (s *service) Do(c echo.Context) error {
	return s.get(c)
}

func (s *service) get(c echo.Context) error {
	rootPath := utils.GetMyRootPath(c)
	discovery := models.DiscoveryDocument{
		Issuer:                rootPath + "/",
		TokenEndpoint:         rootPath + wellknown.OAuth2TokenPath,
		JwksURI:               rootPath + wellknown.WellKnownJWKS,
		RevocationEndpoint:    rootPath + wellknown.OAuth2RevokePath,
		IntrospectionEndpoint: rootPath + wellknown.OAuth2IntrospectPath,
		GrantTypesSupported: []string{
			wellknown.OAuth2GrantType_ClientCredentials,
			wellknown.OAuth2GrantType_RefreshToken,
			wellknown.OAuth2GrantType_TokenExchange,
		},
	}
	return c.JSONPretty(http.StatusOK, discovery, "  ")

}
