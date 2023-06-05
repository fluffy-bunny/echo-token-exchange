package ready

import (
	"context"
	"echo-starter/internal/wellknown"
	"net/http"
	"time"

	contracts_probe "echo-starter/internal/contracts/probe"

	"github.com/catmullet/go-workers"
	contracts_handler "github.com/fluffy-bunny/fluffycore/echo/contracts/handler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	di "github.com/dozm/di"
	"github.com/labstack/echo/v4"
)

type (
	service struct {
		Probes []contracts_probe.IProbe `inject:""`
		runner workers.Runner
	}
	probeWorker struct{}
)

var stemService *service

func init() {
	var _ contracts_handler.IHandler = (*service)(nil)
}
func (s *service) Ctor(probes []contracts_probe.IProbe) (*service, error) {
	obj := &service{
		Probes: probes,
	}
	obj.runner = workers.NewRunner(context.Background(), &probeWorker{}, int64(len(s.Probes))).Start()
	obj.runner.SetTimeout(time.Second * 5)

	return obj, nil
}

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder di.ContainerBuilder) {
	contracts_handler.AddScopedIHandleWithMetadata[*service](builder,
		stemService.Ctor,
		[]contracts_handler.HTTPVERB{
			contracts_handler.GET,
		},
		wellknown.ReadyPath,
	)

}

func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (s *service) Do(c echo.Context) error {
	ctx := c.Request().Context()
	log := zerolog.Ctx(ctx).With().Logger()

	for _, probe := range s.Probes {
		log.Debug().Str("probe", probe.GetName()).Msg("issuing probe")
		s.runner.Send(probe)
	}
	log.Debug().Msg("Waiting for probes to complete")
	err := s.runner.Wait()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "ok")
}
func (s *probeWorker) Work(in interface{}, out chan<- interface{}) error {
	ctx := context.Background()
	ctx = log.Logger.With().Caller().Logger().WithContext(ctx)
	var err error
	if probe, ok := in.(contracts_probe.IProbe); ok {
		log.Debug().Str("probe", probe.GetName()).Msg("probe it")
		err = probe.Probe(ctx)
		if err != nil {
			log.Error().Err(err).Msg("probe failed")
		}
	}
	return err

}
