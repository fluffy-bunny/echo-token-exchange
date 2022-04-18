package models

import "github.com/golang-jwt/jwt"

type (
	Claims  map[string]interface{}
	IClaims interface {
		Valid() error
		Set(key string, value interface{}) error
		Delete(key string) error
		Get(key string) interface{}
		JwtClaims() jwt.Claims
		Claims() Claims
	}
)

// Valid claims verification
func (a *Claims) Valid() error {
	return nil
}
func (a *Claims) Claims() Claims {
	return *a
}

func (a *Claims) Set(key string, value interface{}) error {
	(*a)[key] = value
	return nil
}
func (a *Claims) Delete(key string) error {
	delete(*a, key)
	return nil
}
func (a *Claims) Get(key string) interface{} {
	return (*a)[key]
}
func (a *Claims) JwtClaims() jwt.Claims {
	return a
}
func assertImplementation() {
	var _ IClaims = (*Claims)(nil)
	var _ jwt.Claims = (*Claims)(nil)
}
