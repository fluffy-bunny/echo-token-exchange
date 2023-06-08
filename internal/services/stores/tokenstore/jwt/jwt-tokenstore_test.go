package jwt

import (
	"os"
	"testing"

	"echo-starter/tests"

	contracts_config "echo-starter/internal/contracts/config"
	services_stores_jwttoken "echo-starter/internal/services/stores/jwttoken"
	services_stores_keymaterial "echo-starter/internal/services/stores/keymaterial"
	"echo-starter/internal/services/stores/tokenstore"

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

		builder := di.Builder()
		AddSingletonITokenStore(builder)
		services_stores_jwttoken.AddSingletonIJwtTokenStore(builder)
		services_stores_keymaterial.AddSingletonIKeyMaterial(builder)
		di.AddInstance[*contracts_config.Config](builder, &contracts_config.Config{
			SigningKeys: string(data),
		})

		ctn := builder.Build()
		tokenstore.RunTestSuite(t, ctn)

	})
}
