package TokenExchangeTokenHandler

// https://datatracker.ietf.org/doc/html/draft-ietf-oauth-token-exchange-12#section-2.1
import (
	"context"
	contracts_claimsprovider "echo-starter/internal/contracts/claimsprovider"
	contracts_jwtvalidator "echo-starter/internal/contracts/jwtvalidator"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"errors"

	oauth2_errors "github.com/go-oauth2/oauth2/v4/errors"

	"echo-starter/internal/models"
	"echo-starter/internal/utils"
	"net/http"
	"reflect"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		ClaimsProvider contracts_claimsprovider.IClaimsProvider `inject:""`
		JWTValidator   contracts_jwtvalidator.IJWTValidator     `inject:""`
	}
	validated struct {
		scopes             []string
		subjectToken       string
		subjectTokenType   string
		actorToken         string
		actorTokenType     string
		requestedTokenType string
		audience           string
		resource           string
	}
)

func assertImplementation() {
	var _ contracts_tokenhandlers.ITokenExchangeTokenHandler = (*service)(nil)
}
func init() {
	assertImplementation()
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedITokenExchangeTokenHandler registers the *service.
func AddScopedITokenExchangeTokenHandler(builder *di.Builder) {
	contracts_tokenhandlers.AddScopedITokenExchangeTokenHandler(builder, reflectType)
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
	safeAddParam("subject_token")
	safeAddParam("subject_token_type")
	safeAddParam("actor_token")
	safeAddParam("actor_token_type")
	safeAddParam("requested_token_type")
	safeAddParam("audience")
	safeAddParam("resource")

	if _, ok := validated.Params["subject_token"]; !ok {

		return nil, errors.New("subject_token is required")
	}

	return validated, nil
}
func (s *service) ProcessTokenRequest(ctx context.Context,
	result *contracts_tokenhandlers.ValidatedTokenRequestResult) (models.IClaims, error) {
	claims := make(models.Claims)

	token, err := s.JWTValidator.ValidateJWTRaw(ctx, result.Params["subject_token"])
	if err != nil {
		return nil, oauth2_errors.ErrInvalidRequest
	}
	claims["sub"] = token.Subject()
	for k, v := range token.PrivateClaims() {
		claims[k] = v
	}
	//validated := data.(*validated)
	result.RefreshTokenHandle = "hi"
	return &claims, nil
}
