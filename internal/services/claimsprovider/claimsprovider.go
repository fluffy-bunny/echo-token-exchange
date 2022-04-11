package claimsprovider

import (
	contracts_claimsprovider "echo-starter/internal/contracts/claimsprovider"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/wellknown"
	"errors"
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/golang/mock/gomock"
)

type (
	service struct {
	}
	serviceMock struct {
	}
)

func assertImplementation() {
	var _ contracts_claimsprovider.IClaimsProvider = (*service)(nil)
}

var mockProfileStore map[string]contracts_tokenhandlers.Claims

func init() {
	mockProfileStore = make(map[string]contracts_tokenhandlers.Claims)

	mockProfileStore[""] = make(contracts_tokenhandlers.Claims)
	mockProfileStore["profile1"] = make(contracts_tokenhandlers.Claims)
	mockProfileStore["profile1"][wellknown.ClaimTypeDeep] = []string{
		wellknown.ClaimValueRead,
		wellknown.ClaimValueReadWrite,
		wellknown.ClaimValueReadWriteAll,
	}
	mockProfileStore["profile2"] = make(contracts_tokenhandlers.Claims)
	mockProfileStore["profile2"][wellknown.ClaimTypeDeep] = []string{
		wellknown.ClaimValueRead,
		wellknown.ClaimValueReadWrite,
	}
	mockProfileStore["profile3"] = make(contracts_tokenhandlers.Claims)
	mockProfileStore["profile3"][wellknown.ClaimTypeDeep] = []string{
		wellknown.ClaimValueRead,
	}
}

var reflectType = reflect.TypeOf((*service)(nil))
var reflectTypeMock = reflect.TypeOf((*serviceMock)(nil))

// AddSingletonIClaimsProvider registers the *service as a singleton.
func AddSingletonIClaimsProvider(builder *di.Builder) {
	contracts_claimsprovider.AddSingletonIClaimsProvider(builder, reflectType)
}

func AddSingletonIClaimsProviderMock(builder *di.Builder, ctrl *gomock.Controller) {
	contracts_claimsprovider.AddSingletonIClaimsProvider(builder, reflectTypeMock)
}
func (s *service) Ctor() {}
func (s *service) GetProfiles(userID string) ([]string, error) {
	return []string{"profile1", "profile2", "profile3"}, nil
}
func (s *service) GetClaims(userID string, profile string) (contracts_tokenhandlers.Claims, error) {
	return nil, errors.New("not implemented")
}
func (s *serviceMock) GetProfiles(userID string) ([]string, error) {
	return []string{"profile1", "profile2", "profile3"}, nil
}
func (s *serviceMock) GetClaims(userID string, profile string) (contracts_tokenhandlers.Claims, error) {
	claims, ok := mockProfileStore[profile]
	if !ok {
		return nil, errors.New("profile not found")
	}
	return claims, nil
}
