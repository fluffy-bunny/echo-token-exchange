package inmemory

import (
	contracts_go_oauth2_oauth2 "echo-starter/internal/contracts/go-oauth2/oauth2"
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"

	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
)

func assertImplementation() {
	var _ contracts_go_oauth2_oauth2.IClientStore = (*store.ClientStore)(nil)
}

func AddSingletonIClientStore(builder *di.Builder) {
	reflectType := reflect.TypeOf((*store.ClientStore)(nil))
	contracts_go_oauth2_oauth2.AddSingletonIClientStoreByFunc(builder, reflectType, func(ctn di.Container) (interface{}, error) {

		clientStore := store.NewClientStore()
		clientStore.Set("000000", &models.Client{
			ID:     "000000",
			Secret: "999999",
			Domain: "http://localhost",
		})
		return clientStore, nil
	})
}
