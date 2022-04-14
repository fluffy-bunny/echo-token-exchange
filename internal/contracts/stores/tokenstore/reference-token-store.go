package tokenstore

import (
	"context"
	"time"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IReferenceTokenStore,IInternalReferenceTokenStore"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/stores/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/stores/$GOPACKAGE IReferenceTokenStore,IInternalReferenceTokenStore

type (
	ReferenceTokenInfo struct {
		ClientID   string                 `json:"client_id"`
		Subject    string                 `json:"subject"`
		Response   map[string]interface{} `json:"response"`
		Expiration time.Time              `json:"expiration"`
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
