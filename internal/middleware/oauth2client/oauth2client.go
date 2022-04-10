package claimsprincipal

import (
	"echo-starter/internal/models"

	contracts_clients "echo-starter/internal/contracts/clients"
	"echo-starter/internal/utils"
	"echo-starter/internal/wellknown"

	core_wellknown "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/wellknown"
	oauth2_server "github.com/go-oauth2/oauth2/v4/server"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/labstack/echo/v4"
)

func _clientInfoHandler(c echo.Context, clientStore contracts_clients.IClientStore) (client *models.Client, err error) {
	r := c.Request()
	clientID, clientSecret, err := oauth2_server.ClientBasicHandler(r)
	if err != nil {
		clientID, clientSecret, err = oauth2_server.ClientFormHandler(r)
	}
	if err != nil {
		return nil, err
	}

	client, _, _ = clientStore.GetClient(r.Context(), clientID)
	var match bool
	for _, sc := range client.ClientSecrets {
		match, _ = utils.ComparePasswordHash(clientSecret, sc.Value)
		if match {
			break
		}
	}
	if !match {
		err = errors.ErrInvalidClient
	}
	if err != nil {
		client = nil
	}
	return
}

func AuthenticateOAuth2Client(root di.Container) echo.MiddlewareFunc {

	clientStore := contracts_clients.GetIClientStoreFromContainer(root)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			switch c.Request().URL.Path {
			case wellknown.OAuth2TokenPath:
			case wellknown.OAuth2RevokePath:
			default:
				return next(c)
			}
			client, err := _clientInfoHandler(c, clientStore)
			if err != nil {
				return c.JSON(401, errors.ErrInvalidClient)
			}

			scopedContainer := c.Get(core_wellknown.SCOPED_CONTAINER_KEY).(di.Container)
			clientRequest := contracts_clients.GetIClientRequestInternalFromContainer(scopedContainer)
			clientRequest.SetClient(client)

			return next(c)
		}
	}
}
