package discovery

import (
	"echo-starter/internal/utils"
	"echo-starter/internal/wellknown"
	"net/http"

	models "echo-starter/internal/models"

	di "github.com/dozm/di"
	contracts_handler "github.com/fluffy-bunny/fluffycore/echo/contracts/handler"
	echo "github.com/labstack/echo/v4"
)

type (
	service struct{}
)

var stemService *service = new(service)

func init() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

func (s *service) Ctor() (*service, error) {

	return &service{}, nil
}

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder di.ContainerBuilder) {
	contracts_handler.AddScopedIHandleWithMetadata[*service](builder,
		stemService.Ctor,
		[]contracts_handler.HTTPVERB{
			contracts_handler.GET,
		},
		wellknown.WellKnownOpenIDCOnfiguationPath)

}

func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

// OIDC Discovery godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} models.DiscoveryDocument
// @Router /.well-known/openid-configuration [get]

func (s *service) Do(c echo.Context) error {
	return s.get(c)
}

func (s *service) get(c echo.Context) error {
	rootPath := utils.GetMyRootPath(c)
	discovery := models.DiscoveryDocument{
		Issuer:        rootPath + "/",
		TokenEndpoint: rootPath + wellknown.OAuth2TokenPath,
		JwksURI:       rootPath + wellknown.WellKnownJWKS,
		//	RevocationEndpoint:    rootPath + wellknown.OAuth2RevokePath,
		//	IntrospectionEndpoint: rootPath + wellknown.OAuth2IntrospectPath,
		GrantTypesSupported: []string{
			wellknown.OAuth2GrantType_ClientCredentials,
			wellknown.OAuth2GrantType_RefreshToken,
			wellknown.OAuth2GrantType_TokenExchange,
		},
	}
	return c.JSONPretty(http.StatusOK, discovery, "  ")

}
