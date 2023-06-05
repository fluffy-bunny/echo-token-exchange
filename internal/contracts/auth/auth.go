package auth

import (
	"context"

	"golang.org/x/oauth2"
)

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IInternalTokenStore,ITokenStore

type (

	// ITokenStore is a SCOPED store so nothing global
	ITokenStore interface {
		GetToken(ctx context.Context) (*oauth2.Token, error)
		Clear(ctx context.Context) error
	}
	IInternalTokenStore interface {
		ITokenStore
		GetTokenByIdempotencyKey(ctx context.Context, bindingKey string) (*oauth2.Token, error)
		StoreTokenByIdempotencyKey(ctx context.Context, bindingKey string, token *oauth2.Token) error
		SlideOutExpiration(ctx context.Context) error
	}
)
