package clientrequest

import (
	"reflect"

	contracts_clients "echo-starter/internal/contracts/stores/clients"

	di "github.com/dozm/di"
)

type (
	service struct {
		contracts_clients.CommonClientRequest
	}
)

var stemService *service

func init() {
	var _ contracts_clients.IClientRequest = (*service)(nil)
}
func (s *service) Ctor() (*service, error) {
	return &service{}, nil
}

// AddScopedIClientRequest registers the *service as a scoped.
func AddScopedIClientRequest(builder di.ContainerBuilder) {
	di.AddSingleton[*service](
		builder,
		stemService.Ctor,
		reflect.TypeOf((*contracts_clients.IClientRequest)(nil)),
		reflect.TypeOf((*contracts_clients.IClientRequestInternal)(nil)))

}
