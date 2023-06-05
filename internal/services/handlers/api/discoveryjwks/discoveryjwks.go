package discoveryjwks

import (
	"echo-starter/internal/models"
	"echo-starter/internal/wellknown"
	"net/http"

	contracts_stores_keymaterial "echo-starter/internal/contracts/stores/keymaterial"

	di "github.com/dozm/di"
	contracts_handler "github.com/fluffy-bunny/fluffycore/echo/contracts/handler"
	echo "github.com/labstack/echo/v4"
)

type (
	service struct {
		KeyMaterial contracts_stores_keymaterial.IKeyMaterial `inject:""`
	}
)

var stemService *service

func init() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

func (s *service) Ctor(keyMaterial contracts_stores_keymaterial.IKeyMaterial) (*service, error) {

	return &service{
		KeyMaterial: keyMaterial,
	}, nil
}

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder di.ContainerBuilder) {
	contracts_handler.AddScopedIHandleWithMetadata[*service](builder,
		stemService.Ctor,
		[]contracts_handler.HTTPVERB{
			contracts_handler.GET,
		},
		wellknown.WellKnownJWKS)

}

func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

func (s *service) Do(c echo.Context) error {
	return s.get(c)
}

func (s *service) get(c echo.Context) error {
	type JWKS struct {
		Keys []*models.PublicJwk `json:"keys"`
	}

	keys, _ := s.KeyMaterial.GetPublicWebKeys()
	return c.JSONPretty(http.StatusOK, &JWKS{Keys: keys}, "  ")

}
