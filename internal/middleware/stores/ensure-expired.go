package stores

import (
	"context"

	contracts_stores_refreshtoken "echo-starter/internal/contracts/stores/refreshtoken"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

const middlewareLogName = "ensure-clear-expired-storage-items"

func EnsureClearExpiredStorageItems(container di.Container) echo.MiddlewareFunc {
	referenceTokenStore, _ := contracts_stores_tokenstore.SafeGetIInternalReferenceTokenStoreFromContainer(container)
	refreshTokenStore, _ := contracts_stores_refreshtoken.SafeGetIInternalRefreshTokenStoreFromContainer(container)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if referenceTokenStore != nil {
				referenceTokenStore.RemoveExpired(context.Background())
			}
			if refreshTokenStore != nil {
				refreshTokenStore.RemoveExpired(context.Background())
			}
			return next(c)
		}
	}
}
