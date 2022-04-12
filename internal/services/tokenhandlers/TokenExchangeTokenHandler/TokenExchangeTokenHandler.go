package TokenExchangeTokenHandler

// https://datatracker.ietf.org/doc/html/draft-ietf-oauth-token-exchange-12#section-2.1
import (
	"context"
	"net/http"
	"reflect"

	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/utils"

	contracts_claimsprovider "echo-starter/internal/contracts/claimsprovider"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"

	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		ClaimsProvider contracts_claimsprovider.IClaimsProvider `inject:""`
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
	var _ contracts_tokenhandlers.IClientCredentialsTokenHandler = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedITokenExchangeTokenHandler registers the *service.
func AddScopedITokenExchangeTokenHandler(builder *di.Builder) {
	contracts_tokenhandlers.AddScopedITokenExchangeTokenHandler(builder, reflectType)
}

func (s *service) ValidationTokenRequest(r *http.Request) (result *contracts_tokenhandlers.ValidatedTokenRequestResult, err error) {
	validated := &contracts_tokenhandlers.ValidatedTokenRequestResult{
		GrantType: r.FormValue("grant_type"),
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

	return validated, nil
}
func (s *service) ProcessTokenRequest(ctx context.Context, result *contracts_tokenhandlers.ValidatedTokenRequestResult) (contracts_tokenhandlers.Claims, error) {
	claims := make(contracts_tokenhandlers.Claims)
	//validated := data.(*validated)

	return claims, nil
}
