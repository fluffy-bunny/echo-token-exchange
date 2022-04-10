package oauth2

import (
	"context"
	"echo-starter/internal/models"
	"net/http"
	"time"

	d_oauth2 "github.com/go-oauth2/oauth2/v4"
)

type (
	// GenerateBasic provide the basis of the generated token data
	GenerateBasic struct {
		Client    *models.Client
		UserID    string
		CreateAt  time.Time
		TokenInfo d_oauth2.TokenInfo
		Request   *http.Request
	}

	// AuthorizeGenerate generate the authorization code interface
	AuthorizeGenerate interface {
		Token(ctx context.Context, data *GenerateBasic) (code string, err error)
	}

	// AccessGenerate generate the access and refresh tokens interface
	AccessGenerate interface {
		Token(ctx context.Context, data *GenerateBasic, isGenRefresh bool) (access, refresh string, err error)
	}
)