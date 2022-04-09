package models

const (
	NAMESPACE_NAME  string = "artificer-ns"
	MAX_SUBJECT_LEN int    = 64
	MAX_SCOPE_LEN   int    = 1024
)

type TokenUsage int

const (
	ReUse TokenUsage = iota
	OneTimeOnly
)

type TokenExpiration int

const (
	Sliding TokenExpiration = iota
	Absolute
)

type AccessTokenType int

const (
	Jwt AccessTokenType = iota
	Reference
)

type Claim struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Secret struct {
	Value      string `json:"value"`
	Expiration int    `json:"expiration"`
}

type Client struct {
	Enabled                          bool            `json:"enabled"`
	ClientID                         string          `json:"client_id"`
	ClientName                       string          `json:"client_name"`
	Description                      string          `json:"description"`
	Namespace                        string          `json:"namespace"`
	RequireRefreshClientSecret       bool            `json:"require_refresh_client_secret"`
	AllowOfflineAccess               bool            `json:"allow_offline_access"`
	AccessTokenLifetime              int             `json:"access_token_lifetime"`
	AbsoluteRefreshTokenLifetime     int             `json:"absolute_refresh_token_lifetime"`
	SlidingRefreshTokenLifetime      int             `json:"sliding_refresh_token_lifetime"`
	UpdateAccessTokenClaimsOnRefresh bool            `json:"update_access_token_claims_on_refresh"`
	RefreshTokenUsage                TokenUsage      `json:"refresh_token_usage"`
	RefreshTokenExpiration           TokenExpiration `json:"refresh_token_expiration"`
	AccessTokenType                  AccessTokenType `json:"access_token_type"`
	IncludeJwtId                     bool            `json:"include_jwt_id"`
	Claims                           []Claim         `json:"claims"`
	AlwaysSendClientClaims           bool            `json:"always_send_client_claims"`
	AlwaysIncludeUserClaimsInIdToken bool            `json:"always_include_user_claims_in_id_token"`
	AllowedScopes                    []string        `json:"allowed_scopes"`
	AllowedGrantTypes                []string        `json:"allowed_grant_types"`
	AllowedGrantTypesMap             map[string]interface{}
	RequireClientSecret              bool     `json:"require_client_secret"`
	ClientSecrets                    []Secret `json:"client_secrets"`
}
