package ClientCredentialsTokenHandler

import (
	"context"
	"net/http"
	"reflect"
	"strings"

	contracts_apiresources "echo-starter/internal/contracts/apiresources"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/utils"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		APIResources contracts_apiresources.IAPIResources `inject:""`
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

	return validated, nil
}
func (s *service) ProcessTokenRequest(ctx context.Context, result *contracts_tokenhandlers.ValidatedTokenRequestResult) (contracts_tokenhandlers.Claims, error) {
	claims := make(contracts_tokenhandlers.Claims)

	// the general processor will add all the standard claims.
	// these AUD claims are added because of our apiresources model
	audienceSet := core_hashset.NewStringSet()
	apiResourceScopeSet, _ := s.APIResources.GetApiResourceScopes()
	scopes := strings.Split(result.Params["scope"], " ")
	for _, sc := range scopes {
		if apiResourceScopeSet.Contains(sc) {
			apiResource, _, _ := s.APIResources.GetApiResourceByScope(sc)
			if apiResource != nil {
				audienceSet.Add(apiResource.Name)
			}
		}
	}
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
	return claims, nil
}
