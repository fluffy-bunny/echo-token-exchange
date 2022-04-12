package refreshtoken

import (
	"context"
	"time"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IRefreshTokenStore"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/stores/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/stores/$GOPACKAGE IRefreshTokenStore

type (
	RefreshTokenInfo struct {
		ClientID           string            `json:"client_id"`
		Subject            string            `json:"subject"`
		Scopes             []string          `json:"scopes"`
		GrantType          string            `json:"grant_type"`
		Params             map[string]string `json:"params"`
		Expiration         time.Time         `json:"expiration"`
		AbsoluteExpiration time.Time         `json:"absolute_expiration"`
	}
	IRefreshTokenStore interface {
		StoreRefreshToken(ctx context.Context, info *RefreshTokenInfo) (handle string, err error)
		GetRefreshToken(ctx context.Context, handle string) (*RefreshTokenInfo, error)
		UpdateRefeshToken(ctx context.Context, handle string, info *RefreshTokenInfo) error
		RemoveRefreshToken(ctx context.Context, handle string) error
		RemoveRefreshTokenByClientID(ctx context.Context, clientID string) error
		RemoveRefreshTokenBySubject(ctx context.Context, subject string) error
		RemoveRefreshTokenByClientIdAndSubject(ctx context.Context, clientID string, subject string) error
	}
)
