package TokenExchangeTokenHandler

// https://datatracker.ietf.org/doc/html/draft-ietf-oauth-token-exchange-12#section-2.1
import (
	"context"
	contracts_claimsprovider "echo-starter/internal/contracts/claimsprovider"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/models"
	"echo-starter/internal/utils"
	"net/http"

	di "github.com/dozm/di"
	core_utils "github.com/fluffy-bunny/fluffycore/utils"
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

var stemService *service

func init() {
	var _ contracts_tokenhandlers.IClientCredentialsTokenHandler = (*service)(nil)
}

func (s *service) Ctor(claimsProvider contracts_claimsprovider.IClaimsProvider) (*service, error) {
	return &service{
		ClaimsProvider: claimsProvider,
	}, nil
}

// AddScopedITokenExchangeTokenHandler registers the *service.
func AddScopedITokenExchangeTokenHandler(builder di.ContainerBuilder) {
	di.AddScoped[contracts_tokenhandlers.ITokenExchangeTokenHandler](builder, stemService.Ctor)
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
func (s *service) ProcessTokenRequest(ctx context.Context, result *contracts_tokenhandlers.ValidatedTokenRequestResult) (models.IClaims, error) {
	claims := make(models.Claims)
	//validated := data.(*validated)

	return &claims, nil
}
