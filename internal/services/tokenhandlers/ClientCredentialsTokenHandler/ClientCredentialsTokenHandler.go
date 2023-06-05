package ClientCredentialsTokenHandler

import (
	"context"
	"net/http"
	"strings"

	contracts_stores_apiresources "echo-starter/internal/contracts/stores/apiresources"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/models"
	"echo-starter/internal/utils"

	di "github.com/dozm/di"
	core_hashset "github.com/fluffy-bunny/fluffycore/gods/sets/hashset"
	core_utils "github.com/fluffy-bunny/fluffycore/utils"
)

type (
	service struct {
		APIResources contracts_stores_apiresources.IAPIResources `inject:""`
	}
)

var stemService *service = new(service)

func init() {
	var _ contracts_tokenhandlers.IClientCredentialsTokenHandler = (*service)(nil)
}

func (s *service) Ctor(apiResources contracts_stores_apiresources.IAPIResources) (*service, error) {
	return &service{
		APIResources: apiResources,
	}, nil
}

// AddScopedIClientCredentialsTokenHandler registers the *service.
func AddScopedIClientCredentialsTokenHandler(builder di.ContainerBuilder) {
	di.AddScoped[contracts_tokenhandlers.IClientCredentialsTokenHandler](builder, stemService.Ctor)
}

func (s *service) ValidationTokenRequest(r *http.Request) (result *contracts_tokenhandlers.ValidatedTokenRequestResult, err error) {
	validated := &contracts_tokenhandlers.ValidatedTokenRequestResult{
		GrantType:          r.FormValue("grant_type"),
		Params:             make(map[string]string),
		RefreshTokenHandle: utils.GenerateHandle(),
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
func (s *service) ProcessTokenRequest(ctx context.Context, result *contracts_tokenhandlers.ValidatedTokenRequestResult) (models.IClaims, error) {
	claims := make(models.Claims)

	// the general processor will add all the standard claims.
	// these AUD claims are added because of our apiresources model
	audienceSet := core_hashset.NewStringSet()
	apiResourceScopeSet, _ := s.APIResources.GetApiResourceScopes()
	scope := result.Params["scope"]
	scopes := strings.Split(scope, " ")
	for _, sc := range scopes {
		if apiResourceScopeSet.Contains(sc) {
			apiResource, _, _ := s.APIResources.GetApiResourceByScope(sc)
			if apiResource != nil {
				audienceSet.Add(apiResource.Name)
			}
		}
	}
	claims["aud"] = audienceSet.Values()
	claims["scope"] = scopes
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
	return &claims, nil
}
