package claimsprovider

import (
	contracts_claimsprovider "echo-starter/internal/contracts/claimsprovider"
	"echo-starter/internal/models"
	"echo-starter/internal/wellknown"
	"errors"
	"reflect"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"

	di "github.com/dozm/di"
)

type (
	service struct {
	}
)

var stemService *service

func init() {
	var _ contracts_claimsprovider.IClaimsProvider = (*service)(nil)
}

var mockProfileClaimsStore map[string]models.Claims
var mockUserProfileStore map[string]*core_hashset.StringSet

func init() {
	mockUserProfileStore = make(map[string]*core_hashset.StringSet)
	mockUserProfileStore["user1"] = core_hashset.NewStringSet()
	mockUserProfileStore["user1"].Add("profile1", "profile2", "profile3")
	mockUserProfileStore["user2"] = core_hashset.NewStringSet()
	mockUserProfileStore["user2"].Add("profile1", "profile2", "profile3")
	mockUserProfileStore["user3"] = core_hashset.NewStringSet()
	mockUserProfileStore["user3"].Add("profile1", "profile2", "profile3")

	mockProfileClaimsStore = make(map[string]models.Claims)

	mockProfileClaimsStore[""] = make(models.Claims)
	mockProfileClaimsStore["profile1"] = make(models.Claims)
	mockProfileClaimsStore["profile1"][wellknown.ClaimTypeDeep] = []string{
		wellknown.ClaimValueRead,
		wellknown.ClaimValueReadWrite,
		wellknown.ClaimValueReadWriteAll,
	}
	mockProfileClaimsStore["profile2"] = make(models.Claims)
	mockProfileClaimsStore["profile2"][wellknown.ClaimTypeDeep] = []string{
		wellknown.ClaimValueRead,
		wellknown.ClaimValueReadWrite,
	}
	mockProfileClaimsStore["profile3"] = make(models.Claims)
	mockProfileClaimsStore["profile3"][wellknown.ClaimTypeDeep] = []string{
		wellknown.ClaimValueRead,
	}
}

func (s *service) Ctor() (*service, error) {
	return &service{}, nil
}

// AddSingletonIClaimsProvider registers the *service as a singleton.
func AddSingletonIClaimsProvider(builder di.ContainerBuilder) {
	di.AddSingleton[*service](builder,
		stemService.Ctor,
		reflect.TypeOf((*contracts_claimsprovider.IClaimsProvider)(nil)),
	)
}

func (s *service) GetProfiles(userID string) (*core_hashset.StringSet, error) {
	return mockUserProfileStore[userID], nil
}
func (s *service) GetClaims(userID string, profile string) (models.IClaims, error) {
	userProfiles, ok := mockUserProfileStore[userID]
	if !ok {
		return nil, errors.New("user not found")
	}

	if !userProfiles.Contains(profile) {
		return nil, errors.New("profile not found")
	}
	claims := mockProfileClaimsStore[profile]
	return &claims, nil
}
