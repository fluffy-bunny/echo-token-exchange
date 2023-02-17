package clientrequest

import (
	"reflect"

	contracts_clients "echo-starter/internal/contracts/stores/clients"

	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		contracts_clients.CommonClientRequest
	}
)

func assertImplementation() {
	var _ contracts_clients.IClientRequest = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedIClientRequest registers the *service as a scoped.
func AddScopedIClientRequest(builder *di.Builder) {
	contracts_clients.AddScopedIClientRequest(builder, reflectType, contracts_clients.ReflectTypeIClientRequestInternal)
}
