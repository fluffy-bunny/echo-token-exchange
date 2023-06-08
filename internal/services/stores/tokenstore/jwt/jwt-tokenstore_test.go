package jwt

import (
	"os"
	"testing"

	"echo-starter/tests"

	contracts_config "echo-starter/internal/contracts/config"
	contracts_jwtvalidator "echo-starter/internal/contracts/jwtvalidator"
	services_jwtvalidator "echo-starter/internal/services/jwtvalidator"
	services_stores_jwttoken "echo-starter/internal/services/stores/jwttoken"
	services_stores_keymaterial "echo-starter/internal/services/stores/keymaterial"
	"echo-starter/internal/services/stores/tokenstore"

	fluffycore_services_common "github.com/fluffy-bunny/fluffycore/services/common"

	di "github.com/dozm/di"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	tests.RunTest(t, func(ctrl *gomock.Controller) {
		data, err := os.ReadFile("./signing-keys.json")
		require.NoError(t, err)
		if err == nil {
			t.Setenv("SIGNING_KEYS", string(data))
		}
		boolPtr := func(b bool) *bool {
			return &b
		}
		builder := di.Builder()
		fluffycore_services_common.AddCommonServices(builder)
		services_jwtvalidator.AddSingletonIJwtValidator(builder)
		AddSingletonITokenStore(builder)
		services_stores_jwttoken.AddSingletonIJwtTokenStore(builder)
		services_stores_keymaterial.AddSingletonIKeyMaterial(builder)
		di.AddInstance[*contracts_config.Config](builder, &contracts_config.Config{
			SigningKeys: string(data),
			JWTValidatorOptions: contracts_jwtvalidator.JWTValidatorOptions{
				ClockSkewMinutes:  5,
				ValidateSignature: boolPtr(true),
				ValidateIssuer:    boolPtr(true),
				Issuer:            "http://localhost:1523/",
			},
		})

		ctn := builder.Build()
		tokenstore.RunTestSuite(t, ctn)

	})
}
