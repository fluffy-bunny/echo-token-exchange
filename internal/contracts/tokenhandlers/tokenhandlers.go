package tokenhandlers

import (
	"context"
	"net/http"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ITokenHandler,IClientCredentialsTokenHandler,IRefreshTokenHandler,ITokenExchangeTokenHandler,ITokenHandlerAccessor,IInternalTokenHandlerAccessor"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE ITokenHandler,IClientCredentialsTokenHandler,IRefreshTokenHandler,ITokenExchangeTokenHandler,ITokenHandlerAccessor,IInternalTokenHandlerAccessor

type (
	Claims        map[string]interface{}
	ITokenHandler interface {
		ValidationTokenRequest(r *http.Request) (result interface{}, err error)
		ProcessTokenRequest(ctx context.Context, data interface{}) (Claims, error)
	}

	IClientCredentialsTokenHandler interface {
		ITokenHandler
	}
	IRefreshTokenHandler interface {
		ITokenHandler
	}
	ITokenExchangeTokenHandler interface {
		ITokenHandler
	}
	ITokenHandlerAccessor interface {
		GetGrantType() string
		GetTokenHandler() ITokenHandler
	}
	IInternalTokenHandlerAccessor interface {
		SetGrantType(string)
		SetTokenHandler(ITokenHandler)
	}
	CommonTokenHandlerAccessor struct {
		tokenHandler ITokenHandler
		grantType    string
	}
)

func (s *CommonTokenHandlerAccessor) GetGrantType() string {
	return s.grantType
}
func (s *CommonTokenHandlerAccessor) SetGrantType(v string) {
	s.grantType = v
}
func (s *CommonTokenHandlerAccessor) GetTokenHandler() ITokenHandler {
	return s.tokenHandler
}
func (s *CommonTokenHandlerAccessor) SetTokenHandler(v ITokenHandler) {
	s.tokenHandler = v
}
