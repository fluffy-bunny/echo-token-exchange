package wellknown

import (
	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
)

// reference: https://datatracker.ietf.org/doc/html/draft-ietf-oauth-token-exchange-12
const (
	OAuth2TokenType_Bearer            = "bearer"
	OAuth2TokenType_JWT               = "urn:ietf:params:oauth:token-type:jwt"
	OAuth2TokenType_IDToken           = "urn:ietf:params:oauth:token-type:id_token"
	OAuth2TokenType_RefreshToken      = "urn:ietf:params:oauth:token-type:refresh_token"
	OAuth2TokenType_AccessToken       = "urn:urn:ietf:params:oauth:token-type:access_token"
	OAuth2GrantType_ClientCredentials = "client_credentials"
	OAuth2GrantType_RefreshToken      = "refresh_token"
	OAuth2GrantType_TokenExchange     = "urn:ietf:params:oauth:grant-type:token-exchange"
)

var SupportedGrantTypes *core_hashset.StringSet

func init() {
	SupportedGrantTypes = core_hashset.NewStringSet(
		OAuth2GrantType_ClientCredentials,
		OAuth2GrantType_RefreshToken,
		OAuth2GrantType_TokenExchange)
}
