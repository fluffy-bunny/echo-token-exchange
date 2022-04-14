package revoke

import (
	"echo-starter/internal/wellknown"
	"net/http"
	"reflect"
	"strings"

	contracts_stores_refreshtoken "echo-starter/internal/contracts/stores/refreshtoken"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"

	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

type (
	service struct {
		RefreshTokenStore   contracts_stores_refreshtoken.IRefreshTokenStore `inject:""`
		ReferenceTokenStore contracts_stores_tokenstore.IReferenceTokenStore `inject:""`
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
		wellknown.OAuth2RevokePath)
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
	if core_utils.IsEmptyOrNil(u.TokenTypeHint) {
		return echo.NewHTTPError(http.StatusBadRequest, "token_type_hint is invalid")
	}
	switch u.TokenTypeHint {
	// REFRESH_TOKEN
	case "refresh_token":
		if err := s.RefreshTokenStore.RemoveRefreshToken(ctx, u.Token); err != nil {
			return err
		}
	case "refresh_token:subject":
		if err := s.RefreshTokenStore.RemoveRefreshTokenBySubject(ctx, u.Token); err != nil {
			return err
		}
	case "refresh_token:client_id":
		if err := s.RefreshTokenStore.RemoveRefreshTokenByClientID(ctx, u.Token); err != nil {
			return err
		}
	case "refresh_token:client_id:subject":
		items := strings.Split(u.Token, ":")
		if len(items) != 2 {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}
		if core_utils.IsEmptyOrNil(items[0]) || core_utils.IsEmptyOrNil(items[1]) {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}
		if err := s.RefreshTokenStore.RemoveRefreshTokenByClientIdAndSubject(ctx, items[0], items[1]); err != nil {
			return err
		}
		// ACCESS_TOKEN
	case "access_token":
		if err := s.ReferenceTokenStore.RemoveReferenceToken(ctx, u.Token); err != nil {
			return err
		}
	case "access_token:subject":
		if err := s.ReferenceTokenStore.RemoveReferenceTokenBySubject(ctx, u.Token); err != nil {
			return err
		}
	case "access_token:client_id":
		if err := s.ReferenceTokenStore.RemoveReferenceTokenByClientID(ctx, u.Token); err != nil {
			return err
		}
	case "access_token:client_id:subject":
		items := strings.Split(u.Token, ":")
		if len(items) != 2 {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}
		if core_utils.IsEmptyOrNil(items[0]) || core_utils.IsEmptyOrNil(items[1]) {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}
		if err := s.ReferenceTokenStore.RemoveReferenceTokenByClientIdAndSubject(ctx, items[0], items[1]); err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, "Ok")
}
