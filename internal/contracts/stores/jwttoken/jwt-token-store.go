package jwttoken

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IJwtTokenStore"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/stores/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/stores/$GOPACKAGE IJwtTokenStore
import (
	"context"
	"echo-starter/internal/models"

	"github.com/golang-jwt/jwt"
)

type (
	IJwtTokenStore interface {
		MintToken(ctx context.Context, standardClaims *jwt.StandardClaims, extras models.IClaims) (token string, err error)
	}
)
