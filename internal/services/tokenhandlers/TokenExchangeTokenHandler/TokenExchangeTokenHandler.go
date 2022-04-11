package TokenExchangeTokenHandler

// https://datatracker.ietf.org/doc/html/draft-ietf-oauth-token-exchange-12#section-2.1
import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strings"

	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"

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

func (s *service) ValidationTokenRequest(r *http.Request) (result interface{}, err error) {

	scope := strings.TrimLeft(r.FormValue("scope"), " ")
	scope = strings.TrimRight(scope, " ")
	validate := &validated{
		scopes: strings.Split(scope, " "),
	}
	validate.subjectToken = r.FormValue("subject_token")
	if core_utils.IsEmptyOrNil(validate.subjectToken) {
		return nil, errors.New("subject_token is required")
	}
	validate.subjectTokenType = r.FormValue("subject_token_type")
	if core_utils.IsEmptyOrNil(validate.subjectTokenType) {
		return nil, errors.New("subject_token_type is required")
	}
	validate.actorToken = r.FormValue("actor_token")
	validate.actorTokenType = r.FormValue("actor_token_type")
	validate.requestedTokenType = r.FormValue("requested_token_type")
	validate.audience = r.FormValue("audience")
	validate.resource = r.FormValue("resource")

	return validate, nil
}
func (s *service) ProcessTokenRequest(ctx context.Context, data interface{}) (contracts_tokenhandlers.Claims, error) {
	claims := make(contracts_tokenhandlers.Claims)
	//validated := data.(*validated)

	return claims, nil
}
