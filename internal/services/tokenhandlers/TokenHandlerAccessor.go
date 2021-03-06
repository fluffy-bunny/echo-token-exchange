package tokenhandlers

import (
	"reflect"

	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"

	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		contracts_tokenhandlers.CommonTokenHandlerAccessor
	}
)

func assertImplementation() {
	var _ contracts_tokenhandlers.ITokenHandlerAccessor = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedITokenHandlerAccessor registers the *service.
func AddScopedITokenHandlerAccessor(builder *di.Builder) {
	contracts_tokenhandlers.AddScopedITokenHandlerAccessor(builder, reflectType, contracts_tokenhandlers.ReflectTypeIInternalTokenHandlerAccessor)
}
