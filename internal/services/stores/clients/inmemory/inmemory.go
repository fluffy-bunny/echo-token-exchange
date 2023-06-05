package inmemory

import (
	"context"
	"reflect"

	contracts_clients "echo-starter/internal/contracts/stores/clients"
	"echo-starter/internal/models"

	di "github.com/dozm/di"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
)

type (
	service struct {
		Logger     contracts_logger.ILogger `inject:""`
		clients    []models.Client
		clientsMap map[string]models.Client
	}
)

func assertImplementation() {
	var _ contracts_clients.IClientStore = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIClientStore registers the *service as a singleton.
func AddSingletonIClientStore(builder di.ContainerBuilder, clients []models.Client) {
	contracts_clients.AddSingletonIClientStoreByFunc(builder, reflectType, func(ctn di.Container) (interface{}, error) {

		cMap := make(map[string]models.Client)
		for _, client := range clients {
			cMap[client.ClientID] = client
		}

		obj := &service{
			clients:    clients,
			clientsMap: cMap,
		}
		return obj, nil
	})
}
func (s *service) GetClient(ctx context.Context, id string) (*models.Client, bool, error) {
	client, found := s.clientsMap[id]
	if !found {
		return nil, false, nil
	}
	var copyOfClient models.Client
	copyOfClient = client
	return &copyOfClient, true, nil
}
