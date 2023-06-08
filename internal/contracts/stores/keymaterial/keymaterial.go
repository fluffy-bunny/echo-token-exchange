package keymaterial

import (
	models "echo-starter/internal/models"

	jwk "github.com/lestrrat-go/jwx/v2/jwk"
)

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/stores/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/stores/$GOPACKAGE IKeyMaterial

type (
	// IKeyMaterial ...
	IKeyMaterial interface {
		GetSigningKey() (*models.SigningKey, error)
		GetSigningKeys() ([]*models.SigningKey, error)
		GetPublicWebKeys() ([]*models.PublicJwk, error)
		CreateKeySet() (jwk.Set, error)
	}
)
