package jwttoken

import (
	"context"
	contracts_config "echo-starter/internal/contracts/config"
	contracts_stores_jwttoken "echo-starter/internal/contracts/stores/jwttoken"
	contracts_stores_keymaterial "echo-starter/internal/contracts/stores/keymaterial"
	"echo-starter/internal/models"
	"fmt"
	"reflect"
	"strings"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
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
	var _ contracts_stores_jwttoken.IJwtTokenStore = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIJwtTokenStore registers the *service as a singleton.
func AddSingletonIJwtTokenStore(builder *di.Builder) {
	contracts_stores_jwttoken.AddSingletonIJwtTokenStore(builder, reflectType)
}

func (s *service) MintToken(ctx context.Context, standardClaims *jwt.StandardClaims, extras models.IClaims) (jwtToken string, err error) {
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

	audienceSet := core_hashset.NewStringSet(standardClaims.Audience)
	if !core_utils.IsNil(extras) {
		extraAudInterface := extras.Get("aud")
		switch extraAudInterface.(type) {
		case string:
			audienceSet.Add(extraAudInterface.(string))
		case []string:
			audienceSet.Add(extraAudInterface.([]string)...)
		}
	}
	extras.Set("aud", audienceSet.Values())
	extras.Set("iss", standardClaims.Issuer)
	if !core_utils.IsEmptyOrNil(standardClaims.Subject) {
		extras.Set("sub", standardClaims.Subject)
	}
	if !core_utils.IsEmptyOrNil(standardClaims.Id) {
		extras.Set("jti", standardClaims.Id)
	}
	if !core_utils.IsEmptyOrNil(standardClaims.IssuedAt) {
		extras.Set("iat", standardClaims.IssuedAt)
	}
	if standardClaims.NotBefore > 0 {
		extras.Set("nbf", standardClaims.NotBefore)
	}
	extras.Set("exp", standardClaims.ExpiresAt)

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
