package keystore

import (
	contracts_config "echo-starter/internal/contracts/config"
	contracts_go_oauth2_oauth2 "echo-starter/internal/contracts/go-oauth2/oauth2"
	"echo-starter/internal/models"
	"encoding/json"
	"reflect"

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
func (s *service) GetSigningKeys() ([]*models.SigningKey, error) {
	return s.signingKeys, nil
}
