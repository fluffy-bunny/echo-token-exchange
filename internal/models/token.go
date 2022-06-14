package models

import (
	"encoding/json"
	"time"
)

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
		ClientID                     string            `json:"client_id" mapstructure:"client_id"`
		Subject                      string            `json:"subject" mapstructure:"subject"`
		Scope                        string            `json:"scope" mapstructure:"scope"`
		GrantType                    string            `json:"grant_type" mapstructure:"grant_type"`
		Params                       map[string]string `json:"params" mapstructure:"params"`
		Expiration                   time.Time         `json:"expiration" mapstructure:"expiration"`
		AbsoluteExpiration           time.Time         `json:"absolute_expiration" mapstructure:"absolute_expiration"`
		RefreshTokenGraceEnabled     bool              `json:"refreshTokenGraceEnabled" mapstructure:"refresh_token_grace_enabled"`
		RefreshTokenGraceTTL         time.Duration     `json:"refreshTokenGraceTTL" mapstructure:"refresh_token_grace_ttl"`
		RefreshTokenGraceMaxAttempts int               `json:"refreshTokenGraceMaxAttempts" mapstructure:"refresh_token_grace_max_attempts"`
		RefreshTokenGraceAttempts    int               `json:"refreshTokenGraceAttempts" mapstructure:"refresh_token_grace_attempts"`
	}
)

const (
	TokenTypeRefreshToken                = "refresh_token"
	TokenTypeRefreshTokenSubject         = "refresh_token:subject"
	TokenTypeRefreshTokenClientId        = "refresh_token:client_id"
	TokenTypeRefreshTokenClientIdSubject = "refresh_token:client_id:subject"

	TokenTypeAccessToken                = "access_token"
	TokenTypeAccessTokenSubject         = "access_token:subject"
	TokenTypeAccessTokenClientId        = "access_token:client_id"
	TokenTypeAccessTokenClientIdSubject = "access_token:client_id:subject"

	TokenTypeReferenceToken   = "reference_token"
	TokenTypeSubjectToken     = "subject_token"
	TokenTypeSubjectTokenType = "subject_token_type"
)

func (s *RefreshTokenInfo) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		ClientID                     string            `json:"client_id" mapstructure:"client_id"`
		Subject                      string            `json:"subject" mapstructure:"subject"`
		Scope                        string            `json:"scope" mapstructure:"scope"`
		GrantType                    string            `json:"grant_type" mapstructure:"grant_type"`
		Params                       map[string]string `json:"params" mapstructure:"params"`
		Expiration                   string            `json:"expiration" mapstructure:"expiration"`
		AbsoluteExpiration           string            `json:"absolute_expiration" mapstructure:"absolute_expiration"`
		RefreshTokenGraceEnabled     bool              `json:"refresh_token_grace_enabled" mapstructure:"refresh_token_grace_enabled"`
		RefreshTokenGraceTTL         time.Duration     `json:"refresh_token_grace_ttl" mapstructure:"refresh_token_grace_ttl"`
		RefreshTokenGraceMaxAttempts int               `json:"refresh_token_grace_max_attempts" mapstructure:"refresh_token_grace_max_attempts"`
		RefreshTokenGraceAttempts    int               `json:"refresh_token_grace_attempts" mapstructure:"refresh_token_grace_attempts"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	s.ClientID = tmp.ClientID
	s.Subject = tmp.Subject
	s.Scope = tmp.Scope
	s.GrantType = tmp.GrantType
	s.Params = tmp.Params
	s.Expiration, err = time.Parse(time.RFC3339, tmp.Expiration)
	if err != nil {
		return err
	}
	s.AbsoluteExpiration, err = time.Parse(time.RFC3339, tmp.AbsoluteExpiration)
	if err != nil {
		return err
	}
	s.RefreshTokenGraceEnabled = tmp.RefreshTokenGraceEnabled
	s.RefreshTokenGraceTTL = tmp.RefreshTokenGraceTTL
	s.RefreshTokenGraceMaxAttempts = tmp.RefreshTokenGraceMaxAttempts
	s.RefreshTokenGraceAttempts = tmp.RefreshTokenGraceAttempts
	return nil

}
