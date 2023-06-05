package home

import (
	"echo-starter/internal/wellknown"
	"net/http"

	contracts_config "echo-starter/internal/contracts/config"

	di "github.com/dozm/di"
	fluffycore_contracts_common "github.com/fluffy-bunny/fluffycore/contracts/common"
	fluffycore_echo_contracts_container "github.com/fluffy-bunny/fluffycore/echo/contracts/container"
	contracts_contextaccessor "github.com/fluffy-bunny/fluffycore/echo/contracts/contextaccessor"
	contracts_handler "github.com/fluffy-bunny/fluffycore/echo/contracts/handler"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type (
	service struct {
		// Required and Useful services that the runtime registers
		//---------------------------------------------------------------------------------------------
		ContainerAccessor   fluffycore_echo_contracts_container.ContainerAccessor `inject:""`
		TimeNow             fluffycore_contracts_common.TimeNow                   `inject:""`
		TimeParse           fluffycore_contracts_common.TimeParse                 `inject:""`
		ClaimsPrincipal     fluffycore_contracts_common.IClaimsPrincipal          `inject:""`
		EchoContextAccessor contracts_contextaccessor.IEchoContextAccessor        `inject:""`
		//---------------------------------------------------------------------------------------------

		// internal services
		Config *contracts_config.Config `inject:""`
	}
)

var stemService *service = new(service)

func (s *service) Ctor(
	config *contracts_config.Config,
	containerAccessor fluffycore_echo_contracts_container.ContainerAccessor,
	TimeNow fluffycore_contracts_common.TimeNow,
	TimeParse fluffycore_contracts_common.TimeParse,
	ClaimsPrincipal fluffycore_contracts_common.IClaimsPrincipal,
	EchoContextAccessor contracts_contextaccessor.IEchoContextAccessor) (*service, error) {
	return &service{
		Config:              config,
		ContainerAccessor:   containerAccessor,
		TimeNow:             TimeNow,
		TimeParse:           TimeParse,
		ClaimsPrincipal:     ClaimsPrincipal,
		EchoContextAccessor: EchoContextAccessor,
	}, nil

}
func init() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder di.ContainerBuilder) {
	contracts_handler.AddScopedIHandleWithMetadata[*service](builder,
		stemService.Ctor,
		[]contracts_handler.HTTPVERB{
			contracts_handler.GET,
		},
		wellknown.HomePath)

}

func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}
func (s *service) Do(c echo.Context) error {
	ctx := c.Request().Context()
	log := zerolog.Ctx(ctx).With().Logger()
	log.Info().Str("timeNow", s.TimeNow().String()).Send()
	return c.Render(http.StatusOK, "views/home/index", map[string]interface{}{})
}
