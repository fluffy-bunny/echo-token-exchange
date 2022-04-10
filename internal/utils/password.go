package utils

import (
	"github.com/alexedwards/argon2id"
	passwordGenerator "github.com/theTardigrade/golang-passwordGenerator"
)

func GeneratePasswordHash(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func ComparePasswordHash(password string, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}

func GeneratePassword() string {
	pg := passwordGenerator.New(
		passwordGenerator.Options{
			Len:                     32,
			IncludeUpperCaseLetters: true,
			IncludeLowerCaseLetters: true,
			IncludeDigits:           true,

			ExcludeAmbiguousRunes: true,
			ExcludeRunesList:      []rune{'X', 'x'},
		},
	)

	pass, err := pg.Generate()
	if err != nil {
		pass = "generate_password_error"
	}
	return pass

}
