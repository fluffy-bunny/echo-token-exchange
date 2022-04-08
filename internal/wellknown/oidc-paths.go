package wellknown

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/store"
)

const (
	LoginPath                       = "/login"
	LogoutPath                      = "/logout"
	OIDCCallbackPath                = "/oidc"
	OAuth2CallbackPath              = "/oauth2"
	ErrorPath                       = "/error"
	UnauthorizedPath                = "/unauthorized"
	WellKnownOpenIDCOnfiguationPath = "/.well-known/openid-configuration"
	WellKnownJWKS                   = "/.well-known/jwks"

	OAuth2TokenPath  = "/token"
	OAuth2RevokePath = "/revoke"
)

func F() oauth2.TokenStore {
	store, _ := store.NewMemoryTokenStore()
	return store
}
