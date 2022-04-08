package keystore

import (
	contracts_config "echo-starter/internal/contracts/config"
	contracts_go_oauth2_oauth2 "echo-starter/internal/contracts/go-oauth2/oauth2"
	"echo-starter/internal/models"
	"encoding/json"
	"reflect"
	"time"

	linq "github.com/ahmetb/go-linq"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		Config      *contracts_config.Config `inject:""`
		Logger      contracts_logger.ILogger `inject:""`
		signingKeys []*models.SigningKey
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
}
func (s *service) GetSigningKey() (*models.SigningKey, error) {
	var signingKeys []*models.SigningKey
	now := time.Now()
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
	return signingKeys[len(signingKeys)-1], nil

}

func (s *service) GetPublicWebKeys() ([]*models.PublicJwk, error) {
	var jwks []*models.PublicJwk
	now := time.Now()
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
	return jwks, nil

}
