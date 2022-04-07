package discovery

import (
	"echo-starter/internal/wellknown"
	"fmt"
	"net/http"
	"reflect"

	"echo-starter/internal/models"

	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	di "github.com/fluffy-bunny/sarulabsdi"
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
func AddScopedIHandler(builder *di.Builder) {
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
func getMyRootPath(c echo.Context) string {
	return fmt.Sprintf("%s://%s", c.Scheme(), c.Request().Host)
}

func (s *service) get(c echo.Context) error {
	rootPath := getMyRootPath(c)
	discovery := models.DiscoveryDocument{
		Issuer:             rootPath,
		TokenEndpoint:      rootPath + wellknown.OAuth2TokenPath,
		JwksURI:            rootPath + wellknown.WellKnownJWKS,
		RevocationEndpoint: rootPath + wellknown.OAuth2RevokePath,
		GrantTypesSupported: []string{
			"refresh_token",
			"urn:ietf:params:oauth:grant-type:token-exchange",
			"client_credentials",
		},
	}
	return c.JSON(http.StatusOK, discovery)

}
