package jwtvalidator

import (
	"context"

	jwxt "github.com/lestrrat-go/jwx/jwt"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IJWTValidator"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IJWTValidator

type (
	IJWTValidator interface {
		ValidateJWTRaw(ctx context.Context, token string) (jwxt.Token, error)
	}
)
