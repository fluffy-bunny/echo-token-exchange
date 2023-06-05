package healthz

import (
	"echo-starter/internal/wellknown"
	"net/http"

	contracts_handler "github.com/fluffy-bunny/fluffycore/echo/contracts/handler"

	di "github.com/dozm/di"
	"github.com/labstack/echo/v4"
)

type (
	service struct {
	}
)

var stemService *service = new(service)

func init() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

func (s *service) Ctor() (*service, error) {
	return &service{}, nil
}

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder di.ContainerBuilder) {
	contracts_handler.AddScopedIHandleWithMetadata[*service](builder,
		stemService.Ctor,
		[]contracts_handler.HTTPVERB{
			contracts_handler.GET,
		},
		wellknown.HealthzPath,
	)

}

func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (s *service) Do(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
