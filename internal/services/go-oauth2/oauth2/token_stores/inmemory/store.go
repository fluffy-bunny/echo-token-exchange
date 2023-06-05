package inmemory

import (
	contracts_go_oauth2_oauth2 "echo-starter/internal/contracts/go-oauth2/oauth2"

	di "github.com/dozm/di"

	"github.com/go-oauth2/oauth2/v4/store"
)

func init() {
	var _ contracts_go_oauth2_oauth2.ITokenStore = (*store.TokenStore)(nil)
}

func AddSingletonITokenStore(builder di.ContainerBuilder) {
	di.AddSingleton[contracts_go_oauth2_oauth2.ITokenStore](builder,
		store.NewMemoryTokenStore,
	)
}
