package models

type (
	APIResource struct {
		Name   string   `json:"name"`
		Scopes []string `json:"scopes"`
	}
)
