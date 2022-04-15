package tokenstore

import (
	"context"
	"echo-starter/internal/models"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ITokenStore,IInternalTokenStore"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/stores/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/stores/$GOPACKAGE ITokenStore,IInternalTokenStore

type (
	ITokenStore interface {
		StoreToken(ctx context.Context, info *models.TokenInfo) (handle string, err error)
		GetToken(ctx context.Context, handle string) (*models.TokenInfo, error)
		UpdateToken(ctx context.Context, handle string, info *models.TokenInfo) error
		RemoveToken(ctx context.Context, handle string) error
		RemoveTokenByClientID(ctx context.Context, clientID string) error
		RemoveTokenBySubject(ctx context.Context, subject string) error
		RemoveTokenByClientIdAndSubject(ctx context.Context, clientID string, subject string) error
	}
	IInternalTokenStore interface {
		ITokenStore
		RemoveExpired(ctx context.Context) error
	}
)
