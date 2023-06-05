package inmemory

import (
	contracts_go_oauth2_oauth2 "echo-starter/internal/contracts/go-oauth2/oauth2"
	"reflect"

	di "github.com/dozm/di"

	"github.com/go-oauth2/oauth2/v4/store"
)

func assertImplementation() {
	var _ contracts_go_oauth2_oauth2.ITokenStore = (*store.TokenStore)(nil)
}

func AddSingletonITokenStore(builder di.ContainerBuilder) {
	reflectType := reflect.TypeOf((*store.TokenStore)(nil))
	contracts_go_oauth2_oauth2.AddSingletonITokenStoreByFunc(builder, reflectType, func(ctn di.Container) (interface{}, error) {
		return store.NewMemoryTokenStore()
	})
}
