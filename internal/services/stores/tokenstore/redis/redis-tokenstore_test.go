package redis

import (
	"testing"

	"echo-starter/tests"

	contracts_config "echo-starter/internal/contracts/config"
	"echo-starter/internal/services/stores/tokenstore"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/golang/mock/gomock"
)

func TestStore(t *testing.T) {
	tests.RunTest(t, func(ctrl *gomock.Controller) {
		builder, _ := di.NewBuilder(di.App, di.Request, "transient")
		config := &contracts_config.Config{
			RedisOptionsReferenceTokenStore: contracts_config.RedisOptions{
				Addr:     "localhost:6379",
				Network:  "tcp",
				Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
			},
		}
		di.AddSingletonTypeByObj(builder, config)
		AddSingletonITokenStore(builder)
		ctn := builder.Build()
		tokenstore.RunTestSuite(t, ctn)

	})
}
