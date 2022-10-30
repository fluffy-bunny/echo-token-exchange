package jwtvalidator

import (
	"context"
	"reflect"

	contracts_config "echo-starter/internal/contracts/config"
	contracts_jwtvalidator "echo-starter/internal/contracts/jwtvalidator"

	jwxt "github.com/lestrrat-go/jwx/jwt"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/auth/oauth2"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		Logger        contracts_logger.ILogger `inject:""`
		Config        *contracts_config.Config `inject:""`
		ouath2Context *oauth2.OAuth2Context
		jwtValidator  *oauth2.JWTValidator
	}
)

func assertImplementation() {
	var _ contracts_jwtvalidator.IJWTValidator = (*service)(nil)
}
func init() {
	assertImplementation()
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIJWTValidator registers the *service as a singleton.
func AddSingletonIJWTValidator(builder *di.Builder) {
	contracts_jwtvalidator.AddSingletonIJWTValidator(builder, reflectType)
}
func (s *service) Ctor() {
	ouath2Context, err := oauth2.BuildOAuth2Context(s.Config.OAuth2Issuer,
		s.Config.OAuth2JWKSUrl, &oauth2.GrpcFuncAuthConfig{
			ClockSkewMinutes: s.Config.JWTValidatorAcceptableClockSkewMinutes,
		})
	if err != nil {
		panic(err)
	}
	s.jwtValidator = oauth2.NewJWTValidator(&oauth2.JWTValidatorOptions{
		OAuth2Document:    ouath2Context.OAuth2Document,
		ClockSkewMinutes:  s.Config.JWTValidatorAcceptableClockSkewMinutes,
		ValidateSignature: &(s.Config.JWTValidatorValidateSignature),
	})

	s.ouath2Context = ouath2Context
}
func (s *service) ValidateJWTRaw(ctx context.Context, token string) (jwxt.Token, error) {
	return s.jwtValidator.ParseTokenRaw(ctx, token)
}
