package ClientCredentialsTokenHandler

import (
	"context"
	"net/http"
	"reflect"
	"strings"

	contracts_apiresources "echo-starter/internal/contracts/apiresources"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		APIResources contracts_apiresources.IAPIResources `inject:""`
		contracts_tokenhandlers.CommonTokenHandlerAccessor
	}
	validated struct {
		scopes []string
	}
)

func assertImplementation() {
	var _ contracts_tokenhandlers.IClientCredentialsTokenHandler = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedIClientCredentialsTokenHandler registers the *service.
func AddScopedIClientCredentialsTokenHandler(builder *di.Builder) {
	contracts_tokenhandlers.AddScopedIClientCredentialsTokenHandler(builder, reflectType)
}

func (s *service) ValidationTokenRequest(r *http.Request) (result interface{}, err error) {

	scope := strings.TrimLeft(r.FormValue("scope"), " ")
	scope = strings.TrimRight(scope, " ")
	validate := &validated{
		scopes: strings.Split(scope, " "),
	}

	return validate, nil
}
func (s *service) ProcessTokenRequest(ctx context.Context, data interface{}) (contracts_tokenhandlers.Claims, error) {
	claims := make(contracts_tokenhandlers.Claims)
	validated := data.(*validated)

	// the general processor will add all the standard claims.
	// these AUD claims are added because of our apiresources model
	audienceSet := core_hashset.NewStringSet()
	apiResourceScopeSet, _ := s.APIResources.GetApiResourceScopes()
	for _, sc := range validated.scopes {
		if apiResourceScopeSet.Contains(sc) {
			apiResource, _, _ := s.APIResources.GetApiResourceByScope(sc)
			if apiResource != nil {
				audienceSet.Add(apiResource.Name)
			}
		}
	}
	if validated != nil {
		claims["aud"] = audienceSet.Values()
		/*
			// tests
			claims["basic"] = true
			claims["basic2"] = []string{"basic2"}
			claims["basic3"] = []string{"2", "3"}

			type basic struct {
				Basic string `json:"basic"`
				Count int    `json:"count"`
			}
			claims["basic4"] = []basic{
				{
					Basic: "basic4",
					Count: 4,
				},
			}
			claims["basic5"] = basic{
				Basic: "basic5",
				Count: 5,
			}
		*/
	}
	return claims, nil
}
