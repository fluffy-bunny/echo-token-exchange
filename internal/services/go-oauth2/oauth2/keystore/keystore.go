package keystore

import (
	contracts_config "echo-starter/internal/contracts/config"
	contracts_go_oauth2_oauth2 "echo-starter/internal/contracts/go-oauth2/oauth2"
	"echo-starter/internal/models"
	"echo-starter/internal/utils/ecdsa"
	"encoding/json"
	"reflect"
	"time"

	linq "github.com/ahmetb/go-linq"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		Config        *contracts_config.Config `inject:""`
		Logger        contracts_logger.ILogger `inject:""`
		signingKeys   []*models.SigningKey
		nextFetchTime time.Time
		signingKey    *models.SigningKey
		jwks          []*models.PublicJwk
	}
)

func assertImplementation() {
	var _ contracts_go_oauth2_oauth2.ISigningKeyStore = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonISigningKeyStore registers the *service as a singleton.
func AddSingletonISigningKeyStore(builder *di.Builder) {
	contracts_go_oauth2_oauth2.AddSingletonISigningKeyStore(builder, reflectType)
}
func (s *service) Ctor() {
	var signingKeys []*models.SigningKey
	err := json.Unmarshal([]byte(s.Config.SigningKeys), &signingKeys)
	if err != nil {
		panic(err)
	}
	s.signingKeys = signingKeys
	s.nextFetchTime = time.Now().Add(time.Duration(-24) * time.Hour)
}
func (s *service) _reloadKeys() {
	now := time.Now()
	if now.After(s.nextFetchTime) {
		var signingKeys []*models.SigningKey
		linq.From(s.signingKeys).Where(func(c interface{}) bool {
			signingKey := c.(*models.SigningKey)
			if now.After(signingKey.NotBefore) && now.Before(signingKey.NotAfter) {
				return true
			}
			return false
		}).Select(func(c interface{}) interface{} {
			signingKey := c.(*models.SigningKey)
			return signingKey
		}).ToSlice(&signingKeys)
		// return the last one.
		s.signingKey = signingKeys[len(signingKeys)-1]

		// strip off the encryption and store the open key for downstream ease of use
		privateKey, publicKey, err := ecdsa.DecodePrivatePem(s.signingKey.Password, s.signingKey.PrivateKey)
		if err != nil {
			panic(err)
		}
		encPriv, _, err := ecdsa.Encode("", privateKey, publicKey)
		s.signingKey.PrivateKey = encPriv

		var jwks []*models.PublicJwk
		linq.From(s.signingKeys).Where(func(c interface{}) bool {
			signingKey := c.(*models.SigningKey)
			if now.After(signingKey.NotBefore) && now.Before(signingKey.NotAfter) {
				return true
			}
			return false
		}).Select(func(c interface{}) interface{} {
			signingKey := c.(*models.SigningKey)
			return &signingKey.PublicJwk
		}).ToSlice(&jwks)
		s.jwks = jwks
	}
}

func (s *service) GetSigningKey() (*models.SigningKey, error) {
	s._reloadKeys()
	return s.signingKey, nil
}

func (s *service) GetPublicWebKeys() ([]*models.PublicJwk, error) {
	s._reloadKeys()
	return s.jwks, nil
}
