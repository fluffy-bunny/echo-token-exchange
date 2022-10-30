package startup

import (
	"context"
	contracts_config "echo-starter/internal/contracts/config"
	"fmt"
	"time"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/auth/oauth2"
)

func doWork(config *contracts_config.Config, stop, stopped chan struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// Do cleaning job
				fmt.Println("Gratefully terminated")
				close(stopped)
				return
			default:

				oauth2.BuildOAuth2Context(config.OAuth2Issuer, config.OAuth2JWKSUrl, nil)
				fmt.Println("Performing work")
				// ...
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	<-stop // Block until stop is closed
	cancel()
}
