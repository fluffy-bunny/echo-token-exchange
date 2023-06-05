package keymaterial

import (
	contracts_config "echo-starter/internal/contracts/config"
	contracts_stores_keymaterial "echo-starter/internal/contracts/stores/keymaterial"
	"echo-starter/internal/models"
	"echo-starter/internal/utils/ecdsa"
	"encoding/json"
	"sync"
	"time"

	linq "github.com/ahmetb/go-linq"
	di "github.com/dozm/di"
)

type (
	service struct {
		Config        *contracts_config.Config `inject:""`
		lock          *sync.RWMutex
		signingKeys   []*models.SigningKey
		nextFetchTime time.Time
		signingKey    *models.SigningKey
		jwks          []*models.PublicJwk
	}
)

var stemService *service

func init() {
	var _ contracts_stores_keymaterial.IKeyMaterial = (*service)(nil)
}

func (s *service) Ctor(config *contracts_config.Config) (*service, error) {
	obj := &service{
		Config: config,
	}
	obj.lock = &sync.RWMutex{}
	var signingKeys []*models.SigningKey
	err := json.Unmarshal([]byte(s.Config.SigningKeys), &signingKeys)
	if err != nil {
		panic(err)
	}
	obj.signingKeys = signingKeys
	obj.nextFetchTime = time.Now().Add(time.Duration(-24) * time.Hour)
	return obj, nil
}

// AddSingletonIKeyMaterial registers the *service as a singleton.
func AddSingletonIKeyMaterial(builder di.ContainerBuilder) {
	di.AddSingleton[contracts_stores_keymaterial.IKeyMaterial](builder, stemService.Ctor)
}

func (s *service) _reloadKeys() {
	now := time.Now()
	if now.After(s.nextFetchTime) {
		//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
		s.lock.Lock()
		defer s.lock.Unlock()
		//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
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
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.RLock()
	defer s.lock.RUnlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--

	return s.signingKey, nil
}

func (s *service) GetPublicWebKeys() ([]*models.PublicJwk, error) {
	s._reloadKeys()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.RLock()
	defer s.lock.RUnlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	return s.jwks, nil
}
