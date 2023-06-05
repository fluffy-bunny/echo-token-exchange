package stores

import (
	"context"

	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"

	di "github.com/dozm/di"
	"github.com/labstack/echo/v4"
)

func EnsureClearExpiredStorageItems(container di.Container) echo.MiddlewareFunc {
	tokenStore := di.Get[contracts_stores_tokenstore.IInternalTokenStore](container)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if tokenStore != nil {
				tokenStore.RemoveExpired(context.Background())
			}
			return next(c)
		}
	}
}
