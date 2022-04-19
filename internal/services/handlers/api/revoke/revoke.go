package revoke

import (
	"echo-starter/internal/models"
	"echo-starter/internal/wellknown"
	"net/http"
	"reflect"
	"strings"

	contracts_background_tasks_removetokens "echo-starter/internal/contracts/background/tasks/removetokens"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"

	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

type (
	service struct {
		ReferenceTokenStore       contracts_stores_tokenstore.ITokenStore                            `inject:""`
		RemoveTokensSingletonTask contracts_background_tasks_removetokens.IRemoveTokensSingletonTask `inject:""`
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
	case models.TokenTypeRefreshToken:
		if err := s.ReferenceTokenStore.RemoveToken(ctx, u.Token); err != nil {
			return err
		}
	case models.TokenTypeRefreshTokenSubject:
		_, err := s.RemoveTokensSingletonTask.EnqueTaskTypeRemoveTokenBySubject(&contracts_background_tasks_removetokens.TokenRemoveBySubject{
			Subject: u.Token,
		})
		if err != nil {
			return err
		}
		/*
		   if err := s.ReferenceTokenStore.RemoveTokenBySubject(ctx, u.Token); err != nil {
		   			return err
		   		}
		*/
	case models.TokenTypeRefreshTokenClientId:
		_, err := s.RemoveTokensSingletonTask.EnqueTaskTokenRemoveByClientID(&contracts_background_tasks_removetokens.TokenRemoveByClientID{
			ClientID: u.Token,
		})
		if err != nil {
			return err
		}
		/*
			if err := s.ReferenceTokenStore.RemoveTokenByClientID(ctx, u.Token); err != nil {
				return err
			}
		*/
	case models.TokenTypeRefreshTokenClientIdSubject:
		items := strings.Split(u.Token, ":")
		if len(items) != 2 {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}
		if core_utils.IsEmptyOrNil(items[0]) || core_utils.IsEmptyOrNil(items[1]) {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}
		_, err := s.RemoveTokensSingletonTask.EnqueTaskTokenRemoveByClientIDAndSubject(&contracts_background_tasks_removetokens.TokenRemoveByClientIDAndSubject{
			TokenRemoveByClientID: contracts_background_tasks_removetokens.TokenRemoveByClientID{
				ClientID: items[0],
			},
			TokenRemoveBySubject: contracts_background_tasks_removetokens.TokenRemoveBySubject{
				Subject: items[1],
			},
		})
		if err != nil {
			return err
		}
		/*
			if err := s.ReferenceTokenStore.RemoveTokenByClientIdAndSubject(ctx, items[0], items[1]); err != nil {
				return err
			}
		*/
		// ACCESS_TOKEN
	case models.TokenTypeAccessToken:
		if err := s.ReferenceTokenStore.RemoveToken(ctx, u.Token); err != nil {
			return err
		}
	case models.TokenTypeAccessTokenSubject:
		if err := s.ReferenceTokenStore.RemoveTokenBySubject(ctx, u.Token); err != nil {
			return err
		}
	case models.TokenTypeAccessTokenClientId:
		if err := s.ReferenceTokenStore.RemoveTokenByClientID(ctx, u.Token); err != nil {
			return err
		}
	case models.TokenTypeAccessTokenClientIdSubject:
		items := strings.Split(u.Token, ":")
		if len(items) != 2 {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}
		if core_utils.IsEmptyOrNil(items[0]) || core_utils.IsEmptyOrNil(items[1]) {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}
		if err := s.ReferenceTokenStore.RemoveTokenByClientIdAndSubject(ctx, items[0], items[1]); err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, "Ok")
}
