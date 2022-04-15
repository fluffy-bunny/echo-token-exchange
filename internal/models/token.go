package models

import "time"

type (
	TokenMetadata struct {
		Type       string    `json:"type" mapstructure:"type"` // refresh_token, reference_token
		ClientID   string    `json:"client_id" mapstructure:"client_id"`
		Subject    string    `json:"subject" mapstructure:"subject"`
		Expiration time.Time `json:"expiration" mapstructure:"expiration"`
		IssedAt    time.Time `json:"issued_at" mapstructure:"issued_at"`
	}

	TokenInfo struct {
		Metadata TokenMetadata          `json:"metadata" mapstructure:"metadata"`
		Data     map[string]interface{} `json:"data" mapstructure:"data"`
	}
	RefreshTokenInfo struct {
		ClientID           string            `json:"client_id" mapstructure:"client_id"`
		Subject            string            `json:"subject" mapstructure:"subject"`
		Scope              string            `json:"scope" mapstructure:"scope"`
		GrantType          string            `json:"grant_type" mapstructure:"grant_type"`
		Params             map[string]string `json:"params" mapstructure:"params"`
		Expiration         time.Time         `json:"expiration" mapstructure:"expiration"`
		AbsoluteExpiration time.Time         `json:"absolute_expiration" mapstructure:"absolute_expiration"`
	}
)
