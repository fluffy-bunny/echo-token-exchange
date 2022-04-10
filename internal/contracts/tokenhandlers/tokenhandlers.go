package tokenhandlers

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IClientCredentialsHandler,ITokenExchangeHandler"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IClientCredentialsHandler,ITokenExchangeHandler

type (
	// IClientCredentialsHandler ...
	IClientCredentialsHandler interface {
	}
	// ITokenExchangeHandler ...
	ITokenExchangeHandler interface {
	}
)
