package jwttoken

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IJwtTokenStore"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/stores/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/stores/$GOPACKAGE IJwtTokenStore
import (
	"github.com/golang-jwt/jwt"
)

type (
	IJwtTokenStore interface {
		MintToken(data jwt.Claims) (token string, err error)
	}
)
