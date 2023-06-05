package jwttoken

import (
	"context"
	contracts_config "echo-starter/internal/contracts/config"
	contracts_stores_keymaterial "echo-starter/internal/contracts/stores/keymaterial"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	"echo-starter/internal/models"
	"fmt"
	"reflect"
	"strings"

	di "github.com/dozm/di"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	"github.com/golang-jwt/jwt"
)

type (
	service struct {
		Config      *contracts_config.Config                  `inject:""`
		Logger      contracts_logger.ILogger                  `inject:""`
		KeyMaterial contracts_stores_keymaterial.IKeyMaterial `inject:""`
	}
)

func assertImplementation() {
	var _ contracts_stores_tokenstore.IJwtTokenStore = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIJwtTokenStore registers the *service as a singleton.
func AddSingletonIJwtTokenStore(builder di.ContainerBuilder) {
	contracts_stores_tokenstore.AddSingletonIJwtTokenStore(builder, reflectType)
}

func (s *service) MintToken(ctx context.Context, extras models.IClaims) (jwtToken string, err error) {
	signingKey, err := s.KeyMaterial.GetSigningKey()
	if err != nil {
		return "", err
	}
	var method jwt.SigningMethod
	switch signingKey.PrivateJwk.Alg {
	case "RS256":
		method = jwt.SigningMethodRS256
	case "RS384":
		method = jwt.SigningMethodRS384
	case "RS512":
		method = jwt.SigningMethodRS512
	case "ES256":
		method = jwt.SigningMethodES256
	case "ES384":
		method = jwt.SigningMethodES384
	case "ES512":
		method = jwt.SigningMethodES512
	default:
		return "", fmt.Errorf("unsupported signing method: %s", signingKey.PrivateJwk.Alg)
	}
	kid := signingKey.Kid
	signedKey := []byte(signingKey.PrivateKey)

	var getKey = func() (interface{}, error) {
		var key interface{}
		if strings.HasPrefix(signingKey.PrivateJwk.Alg, "ES") {
			v, err := jwt.ParseECPrivateKeyFromPEM(signedKey)
			if err != nil {
				return "", err
			}
			key = v
			return key, nil
		}

		v, err := jwt.ParseRSAPrivateKeyFromPEM(signedKey)
		if err != nil {
			return "", err
		}
		key = v
		return key, nil
	}

	token := jwt.NewWithClaims(method, extras.JwtClaims())
	token.Header["kid"] = kid
	key, err := getKey()
	if err != nil {
		return "", err
	}

	// special case, aud is allowed

	jwtToken, err = token.SignedString(key)
	if err != nil {
		return "", err
	}
	return
}
