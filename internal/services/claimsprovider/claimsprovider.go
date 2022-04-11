package claimsprovider

import (
	contracts_claimsprovider "echo-starter/internal/contracts/claimsprovider"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/wellknown"
	"errors"
	"reflect"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"

	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
	}
)

func assertImplementation() {
	var _ contracts_claimsprovider.IClaimsProvider = (*service)(nil)
}

var mockProfileClaimsStore map[string]contracts_tokenhandlers.Claims
var mockUserProfileStore map[string]*core_hashset.StringSet

func init() {
	mockUserProfileStore = make(map[string]*core_hashset.StringSet)
	mockUserProfileStore["user1"] = core_hashset.NewStringSet()
	mockUserProfileStore["user1"].Add("profile1", "profile2", "profile3")
	mockUserProfileStore["user2"] = core_hashset.NewStringSet()
	mockUserProfileStore["user2"].Add("profile1", "profile2", "profile3")
	mockUserProfileStore["user3"] = core_hashset.NewStringSet()
	mockUserProfileStore["user3"].Add("profile1", "profile2", "profile3")

	mockProfileClaimsStore = make(map[string]contracts_tokenhandlers.Claims)

	mockProfileClaimsStore[""] = make(contracts_tokenhandlers.Claims)
	mockProfileClaimsStore["profile1"] = make(contracts_tokenhandlers.Claims)
	mockProfileClaimsStore["profile1"][wellknown.ClaimTypeDeep] = []string{
		wellknown.ClaimValueRead,
		wellknown.ClaimValueReadWrite,
		wellknown.ClaimValueReadWriteAll,
	}
	mockProfileClaimsStore["profile2"] = make(contracts_tokenhandlers.Claims)
	mockProfileClaimsStore["profile2"][wellknown.ClaimTypeDeep] = []string{
		wellknown.ClaimValueRead,
		wellknown.ClaimValueReadWrite,
	}
	mockProfileClaimsStore["profile3"] = make(contracts_tokenhandlers.Claims)
	mockProfileClaimsStore["profile3"][wellknown.ClaimTypeDeep] = []string{
		wellknown.ClaimValueRead,
	}
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIClaimsProvider registers the *service as a singleton.
func AddSingletonIClaimsProvider(builder *di.Builder) {
	contracts_claimsprovider.AddSingletonIClaimsProvider(builder, reflectType)
}

func (s *service) Ctor() {}
func (s *service) GetProfiles(userID string) (*core_hashset.StringSet, error) {
	return mockUserProfileStore[userID], nil
}
func (s *service) GetClaims(userID string, profile string) (contracts_tokenhandlers.Claims, error) {
	userProfiles, ok := mockUserProfileStore[userID]
	if !ok {
		return nil, errors.New("user not found")
	}

	if !userProfiles.Contains(profile) {
		return nil, errors.New("profile not found")
	}
	return mockProfileClaimsStore[profile], nil
}
