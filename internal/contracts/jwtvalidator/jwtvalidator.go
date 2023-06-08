package jwtvalidator

import (
	"context"

	jwxt "github.com/lestrrat-go/jwx/v2/jwt"
)

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IJwtValidator

type (
	JWTValidatorOptions struct {
		ClockSkewMinutes  int    `json:"clockSkewMinutes" mapstructure:"CLOCK_SKEW_MINUTES"`
		ValidateSignature *bool  `json:"validateSignature" mapstructure:"VALIDATE_SIGNATURE"`
		ValidateIssuer    *bool  `json:"validateIssuer" mapstructure:"VALIDATE_ISSUER"`
		Issuer            string `json:"issuer" mapstructure:"ISSUER"`
	}
	// IJwtValidator is a SCOPED store so nothing global
	IJwtValidator interface {
		ParseTokenRaw(ctx context.Context, accessToken string) (jwxt.Token, error)
	}
)
