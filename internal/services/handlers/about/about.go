package about

import (
	"echo-starter/internal/templates"
	"echo-starter/internal/wellknown"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	contracts_stores_keymaterial "echo-starter/internal/contracts/stores/keymaterial"

	golinq "github.com/ahmetb/go-linq/v3"
	di "github.com/dozm/di"
	fluffycore_contracts_common "github.com/fluffy-bunny/fluffycore/contracts/common"
	contracts_container "github.com/fluffy-bunny/fluffycore/echo/contracts/container"
	contracts_handler "github.com/fluffy-bunny/fluffycore/echo/contracts/handler"
	echo "github.com/labstack/echo/v4"
	zerolog "github.com/rs/zerolog"
)

type (
	service struct {
		ClaimsPrincipal   fluffycore_contracts_common.IClaimsPrincipal `inject:""`
		ContainerAccessor contracts_container.ContainerAccessor        `inject:""`
		HandlerFactory    contracts_handler.IHandlerFactory            `inject:""`
		KeyMaterial       contracts_stores_keymaterial.IKeyMaterial    `inject:""`
	}
)

var stemService *service

func (s *service) Ctor(
	claimsPrincipal fluffycore_contracts_common.IClaimsPrincipal,
	containerAccessor contracts_container.ContainerAccessor,
	handlerFactory contracts_handler.IHandlerFactory,
	keyMaterial contracts_stores_keymaterial.IKeyMaterial) (*service, error) {
	keys, _ := s.KeyMaterial.GetPublicWebKeys()
	fmt.Println(keys)
	return &service{
		ClaimsPrincipal:   claimsPrincipal,
		ContainerAccessor: containerAccessor,
		HandlerFactory:    handlerFactory,
		KeyMaterial:       keyMaterial,
	}, nil
}
func init() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder di.ContainerBuilder) {
	contracts_handler.AddScopedIHandleWithMetadata[*service](builder,
		stemService.Ctor,
		[]contracts_handler.HTTPVERB{
			contracts_handler.GET,
		},
		wellknown.AboutPath)

}

func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}
func (s *service) Do(c echo.Context) error {
	ctx := c.Request().Context()
	log := zerolog.Ctx(ctx).With().Logger()
	ctn := s.ContainerAccessor()
	descriptors := ctn.GetDescriptors()
	log.Info().Msg("about")
	type row struct {
		Verbs string
		Path  string
	}

	var rows []row

	golinq.From(descriptors).Select(func(c interface{}) interface{} {
		descriptor := c.(*di.Descriptor)
		found := false
		for _, serviceType := range descriptor.ImplementedInterfaceTypes {
			if serviceType == reflect.TypeOf((*contracts_handler.IHandler)(nil)).Elem() {
				found = true
				break
			}
		}
		if !found {
			return nil
		}
		metadata := descriptor.Metadata
		path := metadata["path"].(string)
		httpVerbs, _ := metadata["httpVerbs"].([]contracts_handler.HTTPVERB)
		verbBldr := strings.Builder{}

		for idx, verb := range httpVerbs {
			verbBldr.WriteString(verb.String())
			if idx < len(httpVerbs)-1 {
				verbBldr.WriteString(",")
			}
		}
		return row{
			Verbs: verbBldr.String(),
			Path:  path,
		}

	}).OrderBy(func(i interface{}) interface{} {
		return i.(row).Path
	}).ToSlice(&rows)

	return templates.Render(c, s.ClaimsPrincipal, http.StatusOK, "views/about/index", map[string]interface{}{
		"defs": rows,
	})
}
