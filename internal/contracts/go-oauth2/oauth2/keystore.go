package oauth2

import "echo-starter/internal/models"

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ISigningKeyStore"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/go-oauth2/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/go-oauth2/$GOPACKAGE ISigningKeyStore

type (
	// ISigningKeyStore ...
	ISigningKeyStore interface {
		GetSigningKey() (*models.SigningKey, error)
		GetPublicWebKeys() ([]*models.PublicJwk, error)
	}
)
