package keymaterial

import "echo-starter/internal/models"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/stores/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/stores/$GOPACKAGE IKeyMaterial

type (
	// IKeyMaterial ...
	IKeyMaterial interface {
		GetSigningKey() (*models.SigningKey, error)
		GetPublicWebKeys() ([]*models.PublicJwk, error)
	}
)
