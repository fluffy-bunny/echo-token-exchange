package inmemory

import (
	"reflect"

	contracts_clients "echo-starter/internal/contracts/clients"
	"echo-starter/internal/models"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
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
func AddSingletonIClientStore(builder *di.Builder, clients []models.Client) {
	contracts_clients.AddSingletonIClientStoreByFunc(builder, reflectType, func(ctn di.Container) (interface{}, error) {

		obj := &service{
			clients:    clients,
			clientsMap: make(map[string]models.Client),
		}
		for _, client := range clients {
			obj.clientsMap[client.ClientID] = client
		}
		return obj, nil
	})
}
func (s *service) GetClient(id string) (*models.Client, bool, error) {
	client, found := s.clientsMap[id]
	if !found {
		return nil, false, nil
	}
	var copyOfClient models.Client
	copyOfClient = client
	return &copyOfClient, true, nil
}
