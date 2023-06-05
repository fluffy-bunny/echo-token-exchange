package inmemory

import (
	"context"

	contracts_clients "echo-starter/internal/contracts/stores/clients"
	"echo-starter/internal/models"

	di "github.com/dozm/di"
)

type (
	service struct {
		clients    []models.Client
		clientsMap map[string]models.Client
	}
)

var stemService *service

func (s *service) Ctor(clients []models.Client) (*service, error) {
	cMap := make(map[string]models.Client)
	for _, client := range clients {
		cMap[client.ClientID] = client
	}

	obj := &service{
		clients:    clients,
		clientsMap: cMap,
	}
	return obj, nil
}
func init() {
	var _ contracts_clients.IClientStore = (*service)(nil)
}

// AddSingletonIClientStore registers the *service as a singleton.
func AddSingletonIClientStore(builder di.ContainerBuilder, clients []models.Client) {
	obj, err := stemService.Ctor(clients)
	if err != nil {
		panic(err)
	}
	di.AddInstance[contracts_clients.IClientStore](builder,
		obj)

}
func (s *service) GetClient(ctx context.Context, id string) (*models.Client, bool, error) {
	client, found := s.clientsMap[id]
	if !found {
		return nil, false, nil
	}
	copyOfClient := client
	return &copyOfClient, true, nil
}
