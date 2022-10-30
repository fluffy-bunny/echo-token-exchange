package JWTRefreshTokenHandler

import (
	"context"
	"net/http"
	"reflect"

	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/models"
	"echo-starter/internal/utils"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	oauth2_errors "github.com/go-oauth2/oauth2/v4/errors"
)

type (
	service struct {
		ReferenceTokenStore contracts_stores_tokenstore.ITokenStore `inject:""`
	}
)

func assertImplementation() {
	var _ contracts_tokenhandlers.IRefreshTokenHandler = (*service)(nil)
}
func init() {
	assertImplementation()
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedIRefreshTokenHandler registers the *service.
func AddScopedIRefreshTokenHandler(builder *di.Builder) {
	contracts_tokenhandlers.AddScopedIRefreshTokenHandler(builder, reflectType)
}

func (s *service) ValidationTokenRequest(r *http.Request) (result *contracts_tokenhandlers.ValidatedTokenRequestResult, err error) {
	validated := &contracts_tokenhandlers.ValidatedTokenRequestResult{
		GrantType: r.FormValue("grant_type"),
		Params:    make(map[string]string),
	}
	var safeAddParam = func(key string) {
		val := utils.TrimLeftAndRight(r.FormValue(key))
		if !core_utils.IsEmptyOrNil(val) {
			validated.Params[key] = val
		}
	}
	safeAddParam(models.TokenTypeRefreshToken)
	return validated, nil
}

func (s *service) ProcessTokenRequest(ctx context.Context, result *contracts_tokenhandlers.ValidatedTokenRequestResult) (models.IClaims, error) {
	//now := time.Now()
	handle := result.Params[models.TokenTypeRefreshToken]

	rt, err := s.ReferenceTokenStore.GetToken(ctx, handle)
	if err != nil {
		return nil, oauth2_errors.ErrInvalidRequest
	}
	if rt == nil {
		return nil, oauth2_errors.ErrInvalidRequest
	}
	if rt.Metadata.Type != models.TokenTypeRefreshToken {
		return nil, oauth2_errors.ErrInvalidRequest
	}

	claims := make(models.Claims)

	return &claims, nil
}
