package rejson

import (
	"log"
	"os"
	"testing"

	contracts_config "echo-starter/internal/contracts/config"
	"echo-starter/internal/services/stores/tokenstore"
	"echo-starter/tests"

	services_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
)

func TestStore(t *testing.T) {
	appEnv := os.Getenv("APPLICATION_ENVIRONMENT")
	if appEnv != "Development" {
		t.Skip("skipping redis tests")
		return
	}
	tests.RunTest(t, func(ctrl *gomock.Controller) {
		builder, _ := di.NewBuilder(di.App, di.Request, "transient")
		config := &contracts_config.Config{
			RedisOptionsReferenceTokenStore: contracts_config.RedisOptions{
				Addr:     "localhost:6379",
				Network:  "tcp",
				Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
			},
		}
		services_logger.AddSingletonILogger(builder)
		di.AddSingletonTypeByObj(builder, config)
		AddSingletonITokenStore(builder)
		ctn := builder.Build()
		redisOptions := &redis.Options{
			Addr:     config.RedisOptionsReferenceTokenStore.Addr,
			Network:  config.RedisOptionsReferenceTokenStore.Network,
			Password: config.RedisOptionsReferenceTokenStore.Password,
			Username: config.RedisOptionsReferenceTokenStore.Username,
		}
		cli := redis.NewClient(redisOptions)

		//		search := redisearch.New(redisOptions)
		//		createSearchIndexes("echoTokenStoreIdx", search)
		//		defer search.DropIndex(context.Background(), "echoTokenStoreIdx", true)
		defer func() {

			if err := cli.Close(); err != nil {
				log.Fatalf("goredis - failed to communicate to redis-server: %v", err)
			}
		}()
		tokenstore.RunTestSuite(t, ctn)

	})
}
