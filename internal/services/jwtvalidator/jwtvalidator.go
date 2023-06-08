package jwtvalidator

import (
	"context"
	contracts_config "echo-starter/internal/contracts/config"
	contracts_jwtvalidator "echo-starter/internal/contracts/jwtvalidator"
	contracts_stores_keymaterial "echo-starter/internal/contracts/stores/keymaterial"
	"time"

	jwk "github.com/lestrrat-go/jwx/v2/jwk"
	jwxt "github.com/lestrrat-go/jwx/v2/jwt"

	di "github.com/dozm/di"
)

type (
	service struct {
		Config      *contracts_config.Config                  `inject:""`
		KeyMaterial contracts_stores_keymaterial.IKeyMaterial `inject:""`
		keySet      jwk.Set
	}
)

var stemService *service = new(service)

func init() {
	var _ contracts_jwtvalidator.IJwtValidator = (*service)(nil)
}
func (s *service) Ctor(
	config *contracts_config.Config,
	keyMaterial contracts_stores_keymaterial.IKeyMaterial,
) (*service, error) {
	keySet, err := keyMaterial.CreateKeySet()
	if err != nil {
		return nil, err
	}

	return &service{
		Config:      config,
		KeyMaterial: keyMaterial,
		keySet:      keySet,
	}, nil
}

// AddSingletonIJwtValidator registers the *service as a singleton.
func AddSingletonIJwtValidator(builder di.ContainerBuilder) {
	di.AddSingleton[contracts_jwtvalidator.IJwtValidator](builder, stemService.Ctor)
}
func (s *service) shouldValidateSignature() bool {
	if s.Config.JWTValidatorOptions.ValidateSignature == nil {
		return true
	}
	return *s.Config.JWTValidatorOptions.ValidateSignature
}
func (s *service) shouldValidateIssuer() bool {
	if s.Config.JWTValidatorOptions.ValidateIssuer == nil {
		return true
	}
	return *s.Config.JWTValidatorOptions.ValidateIssuer
}

func (s *service) ParseTokenRaw(ctx context.Context, accessToken string) (jwxt.Token, error) {
	// Parse the JWT
	parseOptions := []jwxt.ParseOption{}

	if s.shouldValidateSignature() {
		jwkSet := s.keySet
		parseOptions = append(parseOptions, jwxt.WithKeySet(jwkSet))
	}

	token, err := jwxt.ParseString(accessToken, parseOptions...)
	if err != nil {
		return nil, err
	}

	// This set had a key that worked
	var validationOpts []jwxt.ValidateOption
	if s.shouldValidateIssuer() {
		validationOpts = append(validationOpts, jwxt.WithIssuer(s.Config.JWTValidatorOptions.Issuer))
	}
	// Allow clock skew
	validationOpts = append(validationOpts,
		jwxt.WithAcceptableSkew(time.Minute*time.Duration(s.Config.JWTValidatorOptions.ClockSkewMinutes)))

	opts := validationOpts
	err = jwxt.Validate(token, opts...)
	if err != nil {
		return nil, err
	}
	return token, nil
}
