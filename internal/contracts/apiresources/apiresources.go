package apiresources

import "echo-starter/internal/models"

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IAPIResources"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IAPIResources

type (
	// IAPIResources ...
	IAPIResources interface {
		GetAPIResources() ([]models.APIResource, error)
		GetAPIResource(name string) (*models.APIResource, bool, error)
	}
)
