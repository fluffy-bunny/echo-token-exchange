package tokenstore

import (
	"context"
	"echo-starter/internal/models"
)

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/stores/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/stores/$GOPACKAGE IJwtTokenStore

type (
	IJwtTokenStore interface {
		MintToken(ctx context.Context, claims models.IClaims) (token string, err error)
	}
)
