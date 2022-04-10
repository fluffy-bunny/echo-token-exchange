package tokenhandlers

import (
	"context"
	"net/http"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ITokenHandler"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE ITokenHandler

type (
	Claims        map[string]interface{}
	ITokenHandler interface {
		GetGrantType() string
		ValidationTokenRequest(r *http.Request) (result interface{}, err error)
		ProcessTokenRequest(ctx context.Context, data interface{}) (Claims, error)
	}
	CommonTokenHandler struct {
		grantType string
	}
)

func (s *CommonTokenHandler) GetGrantType() string {
	return s.grantType
}
