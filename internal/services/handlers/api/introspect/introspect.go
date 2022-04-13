package introspect

import (
	"echo-starter/internal/wellknown"
	"net/http"
	"reflect"

	contracts_stores_referencetoken "echo-starter/internal/contracts/stores/referencetoken"

	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

type (
	service struct {
		ReferenceTokenStore contracts_stores_referencetoken.IReferenceTokenStore `inject:""`
	}
)

func assertImplementation() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder *di.Builder) {
	contracts_handler.AddScopedIHandlerEx(builder,
		reflectType,
		[]contracts_handler.HTTPVERB{
			contracts_handler.POST,
		},
		wellknown.OAuth2IntrospectPath)
}

func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

type params struct {
	Token         string `param:"token" query:"token" header:"token" form:"token" json:"token" xml:"token"`
	TokenTypeHint string `param:"token_type_hint" query:"token_type_hint" header:"token_type_hint" form:"token_type_hint" json:"token_type_hint" xml:"token_type_hint"`
}

func (s *service) Do(c echo.Context) error {
	return s.post(c)
}

func (s *service) post(c echo.Context) error {
	ctx := c.Request().Context()
	u := new(params)
	if err := c.Bind(u); err != nil {
		return err
	}
	if core_utils.IsEmptyOrNil(u.Token) {
		return echo.NewHTTPError(http.StatusBadRequest, "token is invalid")
	}

	tokenInfo, err := s.ReferenceTokenStore.GetReferenceToken(ctx, u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}

	return c.JSON(http.StatusOK, tokenInfo.Response)
}