package oidc

import (
	"reflect"

	contracts_probe "echo-starter/internal/contracts/probe"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

type (
	service struct {
	}
)

func assertImplementation() {
	var _ contracts_probe.IProbe = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIProbe registers the *service as a singleton.
func AddSingletonIProbe(builder *di.Builder) {
	contracts_probe.AddSingletonIProbe(builder, reflectType)
}
func (s *service) GetName() string {
	return "oidc"
}
func (s *service) Probe() error {
	log.Debug().Str("probe", "oidc").Send()

	return nil
}
