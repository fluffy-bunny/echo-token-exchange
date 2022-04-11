package oauth2

import (
	"context"
	contracts_apiresources "echo-starter/internal/contracts/apiresources"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/models"
	"net/http"
	"time"

	d_oauth2 "github.com/go-oauth2/oauth2/v4"
)

type (
	// GenerateBasic provide the basis of the generated token data
	GenerateBasic struct {
		Client       *models.Client
		APIResources contracts_apiresources.IAPIResources
		UserID       string
		CreateAt     time.Time
		TokenInfo    d_oauth2.TokenInfo
		Request      *http.Request
	}

	// AuthorizeGenerate generate the authorization code interface
	AuthorizeGenerate interface {
		Token(ctx context.Context, data *GenerateBasic) (code string, err error)
	}

	// AccessGenerate generate the access and refresh tokens interface
	AccessGenerate interface {
		Token(ctx context.Context, data *GenerateBasic, isGenRefresh bool, claims contracts_tokenhandlers.Claims) (access, refresh string, err error)
	}
)
