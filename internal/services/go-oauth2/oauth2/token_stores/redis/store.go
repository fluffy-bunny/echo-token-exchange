package inmemory

import (
	contracts_config "echo-starter/internal/contracts/config"
	contracts_go_oauth2_oauth2 "echo-starter/internal/contracts/go-oauth2/oauth2"
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
)

func assertImplementation() {
	var _ contracts_go_oauth2_oauth2.ITokenStore = (*store.TokenStore)(nil)
}

func AddSingletonITokenStore(builder *di.Builder) {
	reflectType := reflect.TypeOf((*store.TokenStore)(nil))
	contracts_go_oauth2_oauth2.AddSingletonITokenStoreByFunc(builder, reflectType, func(ctn di.Container) (interface{}, error) {

		config := ctn.GetByType(contracts_config.ReflectConfigType).(*contracts_config.Config)
		store := oredis.NewRedisStore(&redis.Options{
			Addr:     config.RedisUrl,
			Password: config.RedisPassword,
		})
		return store, nil
	})
}
