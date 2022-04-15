package RefreshTokenHandler

import (
	"context"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/models"
	"echo-starter/internal/utils"
	"echo-starter/internal/wellknown"
	"net/http"
	"reflect"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/mitchellh/mapstructure"
)

type (
	service struct {
		TokenExchangeTokenHandler     contracts_tokenhandlers.ITokenExchangeTokenHandler     `inject:""`
		ClientCredentialsTokenHandler contracts_tokenhandlers.IClientCredentialsTokenHandler `inject:""`
		RefreshTokenStore             contracts_stores_tokenstore.ITokenStore                `inject:""`
	}
	validated struct {
		scopes []string
	}
)

func assertImplementation() {
	var _ contracts_tokenhandlers.IClientCredentialsTokenHandler = (*service)(nil)
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
	safeAddParam("scope")
	safeAddParam("refresh_token")

	return validated, nil
}
func (s *service) ProcessTokenRequest(ctx context.Context, result *contracts_tokenhandlers.ValidatedTokenRequestResult) (models.IClaims, error) {
	handle, _ := result.Params["refresh_token"]
	rt, err := s.RefreshTokenStore.GetToken(ctx, handle)
	if err != nil {
		return nil, errors.ErrInvalidRequest
	}
	if rt == nil {
		return nil, errors.ErrInvalidRequest
	}
	if rt.Metadata.Type != "refresh_token" {
		return nil, errors.ErrInvalidRequest
	}
	if rt.Metadata.ClientID != result.ClientID {
		return nil, errors.New("clientID mismatch")
	}
	// if no scope is passed then we use the scope from the last run
	scope, ok := result.Params["scope"]
	rtInfo := &models.RefreshTokenInfo{}
	err = mapstructure.Decode(rt.Data, rtInfo)
	if err != nil {
		return nil, err
	}
	result.Params = rtInfo.Params
	if ok {
		// override the sone passed into the refresh_token request
		result.Params["scope"] = scope
	}
	newValidatedResult := &contracts_tokenhandlers.ValidatedTokenRequestResult{
		GrantType: rtInfo.GrantType,
		ClientID:  rtInfo.ClientID,
		Params:    result.Params,
	}
	switch rtInfo.GrantType {
	case wellknown.OAuth2GrantType_ClientCredentials:
		return s.ClientCredentialsTokenHandler.ProcessTokenRequest(ctx, newValidatedResult)
	case wellknown.OAuth2GrantType_TokenExchange:
		return s.TokenExchangeTokenHandler.ProcessTokenRequest(ctx, newValidatedResult)

	default:
		return nil, errors.ErrUnsupportedGrantType
	}

}
