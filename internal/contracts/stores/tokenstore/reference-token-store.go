package tokenstore

import (
	"context"
	"echo-starter/internal/models"
	"time"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IReferenceTokenStore,IInternalReferenceTokenStore,ITokenStore"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/stores/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/stores/$GOPACKAGE IReferenceTokenStore,IInternalReferenceTokenStore,ITokenStore

type (
	ReferenceTokenInfo struct {
		ClientID   string                 `json:"client_id"`
		Subject    string                 `json:"subject"`
		Response   map[string]interface{} `json:"response"`
		Expiration time.Time              `json:"expiration"`
	}

	ITokenStore interface {
		StoreToken(ctx context.Context, info *models.TokenInfo) (handle string, err error)
		GetToken(ctx context.Context, handle string) (*models.TokenInfo, error)
		UpdateToken(ctx context.Context, handle string, info *models.TokenInfo) error
		RemoveToken(ctx context.Context, handle string) error
		RemoveTokenByClientID(ctx context.Context, clientID string) error
		RemoveTokenBySubject(ctx context.Context, subject string) error
		RemoveTokenByClientIdAndSubject(ctx context.Context, clientID string, subject string) error
	}
	IReferenceTokenStore interface {
		StoreReferenceToken(ctx context.Context, info *ReferenceTokenInfo) (handle string, err error)
		GetReferenceToken(ctx context.Context, handle string) (*ReferenceTokenInfo, error)
		UpdateReferenceToken(ctx context.Context, handle string, info *ReferenceTokenInfo) error
		RemoveReferenceToken(ctx context.Context, handle string) error
		RemoveReferenceTokenByClientID(ctx context.Context, clientID string) error
		RemoveReferenceTokenBySubject(ctx context.Context, subject string) error
		RemoveReferenceTokenByClientIdAndSubject(ctx context.Context, clientID string, subject string) error
	}
	IInternalReferenceTokenStore interface {
		IReferenceTokenStore
		RemoveExpired(ctx context.Context)
	}
)
