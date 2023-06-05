package clientrequest

import (
	"reflect"

	contracts_clients "echo-starter/internal/contracts/stores/clients"

	di "github.com/dozm/di"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
)

type (
	service struct {
		contracts_clients.CommonClientRequest
		Logger contracts_logger.ILogger `inject:""`
	}
)

func assertImplementation() {
	var _ contracts_clients.IClientRequest = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedIClientRequest registers the *service as a scoped.
func AddScopedIClientRequest(builder di.ContainerBuilder) {
	contracts_clients.AddScopedIClientRequest(builder, reflectType, contracts_clients.ReflectTypeIClientRequestInternal)
}
