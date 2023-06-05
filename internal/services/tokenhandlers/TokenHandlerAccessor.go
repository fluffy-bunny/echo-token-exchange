package tokenhandlers

import (
	"reflect"

	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"

	di "github.com/dozm/di"
)

type (
	service struct {
		contracts_tokenhandlers.CommonTokenHandlerAccessor
	}
)

var stemService *service = new(service)

func init() {
	var _ contracts_tokenhandlers.ITokenHandlerAccessor = (*service)(nil)
	var _ contracts_tokenhandlers.IInternalTokenHandlerAccessor = (*service)(nil)
}
func (s *service) Ctor() (*service, error) {
	return &service{}, nil
}

// AddScopedITokenHandlerAccessor registers the *service.
func AddScopedITokenHandlerAccessor(builder di.ContainerBuilder) {
	di.AddScoped[*service](builder, stemService.Ctor,
		reflect.TypeOf((*contracts_tokenhandlers.ITokenHandlerAccessor)(nil)),
		reflect.TypeOf((*contracts_tokenhandlers.IInternalTokenHandlerAccessor)(nil)),
	)
}
