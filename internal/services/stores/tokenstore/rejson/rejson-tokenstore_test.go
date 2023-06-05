package rejson

import (
	"context"
	"log"
	"os"
	"testing"

	contracts_config "echo-starter/internal/contracts/config"
	"echo-starter/internal/services/stores/tokenstore"
	"echo-starter/tests"

	"github.com/fluffy-bunny/go-redis-search/ftsearch"

	di "github.com/dozm/di"
	services_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
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
			RedisOptions: contracts_config.RedisOptions{
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
			Addr:     config.RedisOptions.Addr,
			Network:  config.RedisOptions.Network,
			Password: config.RedisOptions.Password,
			Username: config.RedisOptions.Username,
		}
		cli := redis.NewClient(redisOptions)
		indexName := "echoTokenStoreIdx"
		var ftSearch *ftsearch.Client
		ftSearch = ftsearch.NewClient(cli)
		results, err := ftSearch.DropIndex(context.Background(), ftsearch.NewDropIndex().WithIndex(indexName))
		if err != nil {
			log.Println(err)
		}
		log.Println(results)

		create := ftsearch.NewCreate().WithIndex(indexName).OnJSON().
			WithSchema(ftsearch.NewSchema().
				WithIdentifier("$.metadata.type").AsAttribute("type").AttributeType("TEXT")).
			WithSchema(ftsearch.NewSchema().
				WithIdentifier("$.metadata.client_id").AsAttribute("client_id").AttributeType("TEXT")).
			WithSchema(ftsearch.NewSchema().
				WithIdentifier("$.metadata.subject").AsAttribute("subject").AttributeType("TEXT"))
		results2, err := ftSearch.CreateIndex(context.Background(), create)
		if err != nil {
			log.Println(err)
		}
		log.Println(results2)
		//		search := redisearch.New(redisOptions)
		//		createSearchIndexes("echoTokenStoreIdx", search)
		//		defer search.DropIndex(context.Background(), "echoTokenStoreIdx", true)
		defer func() {
			// drop the index on the way out
			ftSearch.DropIndex(context.Background(), ftsearch.NewDropIndex().WithIndex(indexName))
			if err := cli.Close(); err != nil {
				log.Fatalf("goredis - failed to communicate to redis-server: %v", err)
			}
		}()
		tokenstore.RunTestSuite(t, ctn)

	})
}
