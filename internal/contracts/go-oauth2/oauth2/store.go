package oauth2

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ITokenStore,IClientStore"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/go-oauth2/$GOPACKAGE ITokenStore,IClientStore
import (
	"github.com/go-oauth2/oauth2/v4"
)

type (
	// ITokenStore ...
	ITokenStore interface {
		oauth2.TokenStore
	}
	// IClientStore ...
	IClientStore interface {
		oauth2.ClientStore
	}
)
