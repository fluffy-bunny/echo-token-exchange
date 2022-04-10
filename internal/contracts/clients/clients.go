package clients

import (
	"context"
	"echo-starter/internal/models"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IClientStore"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IClientStore

type (
	// IClientStore ...
	IClientStore interface {
		GetClient(ctx context.Context, clientID string) (*models.Client, bool, error)
	}
)
