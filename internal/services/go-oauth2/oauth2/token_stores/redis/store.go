package inmemory

import (
	contracts_config "echo-starter/internal/contracts/config"
	contracts_go_oauth2_oauth2 "echo-starter/internal/contracts/go-oauth2/oauth2"

	di "github.com/dozm/di"
	"github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
)

func init() {
	var _ contracts_go_oauth2_oauth2.ITokenStore = (*store.TokenStore)(nil)
}

func ctor(config *contracts_config.Config) (contracts_go_oauth2_oauth2.ITokenStore, error) {
	store := oredis.NewRedisStore(&redis.Options{
		Addr:     config.RedisUrl,
		Password: config.RedisPassword,
	})
	return store, nil
}
func AddSingletonITokenStore(builder di.ContainerBuilder) {
	di.AddSingleton[contracts_go_oauth2_oauth2.ITokenStore](builder, ctor)

}
