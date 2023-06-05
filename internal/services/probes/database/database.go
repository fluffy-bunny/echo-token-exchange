package database

import (
	"context"

	contracts_probe "echo-starter/internal/contracts/probe"

	di "github.com/dozm/di"
	zerolog "github.com/rs/zerolog"
)

type (
	service struct {
	}
)

var stemService *service = new(service)

func init() {
	var _ contracts_probe.IProbe = (*service)(nil)
}
func (s *service) Ctor() (*service, error) {
	return &service{}, nil
}

// AddSingletonIProbe registers the *service as a singleton.
func AddSingletonIProbe(builder di.ContainerBuilder) {
	di.AddSingleton[contracts_probe.IProbe](builder, func() (contracts_probe.IProbe, error) {
		return stemService.Ctor()
	})
}
func (s *service) GetName() string {
	return "database"
}
func (s *service) Probe(ctx context.Context) error {

	log := zerolog.Ctx(ctx).With().Logger()
	log.Debug().Str("probe", "database").Send()
	//return errors.New("DataBase is down")
	return nil
}
