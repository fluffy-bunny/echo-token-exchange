package apiresources

import (
	"echo-starter/internal/models"

	core_hashset "github.com/fluffy-bunny/fluffycore/gods/sets/hashset"
)

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/stores/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/stores/$GOPACKAGE IAPIResources

type (
	// IAPIResources ...
	IAPIResources interface {
		GetAPIResources() ([]models.APIResource, error)
		GetAPIResource(name string) (*models.APIResource, bool, error)
		GetApiResourceByScope(scope string) (*models.APIResource, bool, error)
		GetApiResourceScopes() (*core_hashset.StringSet, error)
	}
)
