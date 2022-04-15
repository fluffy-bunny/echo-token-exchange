package stores

import (
	"context"

	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

const middlewareLogName = "ensure-clear-expired-storage-items"

func EnsureClearExpiredStorageItems(container di.Container) echo.MiddlewareFunc {
	tokenStore, _ := contracts_stores_tokenstore.SafeGetIInternalTokenStoreFromContainer(container)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if tokenStore != nil {
				tokenStore.RemoveExpired(context.Background())
			}
			return next(c)
		}
	}
}
