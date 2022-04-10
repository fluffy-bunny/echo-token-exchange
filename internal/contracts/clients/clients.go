package clients

import (
	"context"
	"echo-starter/internal/models"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IClientStore,IClientTokenRequest,IClientRequest,IClientRequestInternal,IClientTokenRequestInternal"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IClientStore,IClientTokenRequest,IClientRequest,IClientRequestInternal,IClientTokenRequestInternal

type (
	// IClientStore ...
	IClientStore interface {
		GetClient(ctx context.Context, clientID string) (*models.Client, bool, error)
	}

	IClientRequest interface {
		GetClient() *models.Client
	}
	IClientRequestInternal interface {
		IClientRequest
		SetClient(client *models.Client)
	}
	IClientTokenRequest interface {
		IClientRequest
		GetGrantType() string
	}
	IClientTokenRequestInternal interface {
		IClientRequestInternal
		IClientTokenRequest
	}

	CommonClientRequest struct {
		client *models.Client
	}
)

func (s *CommonClientRequest) GetClient() *models.Client {
	return s.client
}
func (s *CommonClientRequest) SetClient(client *models.Client) {
	s.client = client
}
