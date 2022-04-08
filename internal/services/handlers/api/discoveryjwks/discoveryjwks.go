package discoveryjwks

import (
	"echo-starter/internal/wellknown"
	"fmt"
	"net/http"
	"reflect"

	contracts_go_oauth2_oauth2 "echo-starter/internal/contracts/go-oauth2/oauth2"

	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

type (
	service struct {
		SigningKeyStore contracts_go_oauth2_oauth2.ISigningKeyStore `inject:""`
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
			contracts_handler.GET,
		},
		wellknown.WellKnownJWKS)
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

	keys, _ := s.SigningKeyStore.GetPublicWebKeys()
	return c.JSON(http.StatusOK, keys)

}
