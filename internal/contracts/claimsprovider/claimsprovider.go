package claimsprovider

import (
	"echo-starter/internal/models"

	core_hashset "github.com/fluffy-bunny/fluffycore/gods/sets/hashset"
)

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IClaimsProvider

type (
	// IClaimsProvider ...
	IClaimsProvider interface {
		GetProfiles(userID string) (*core_hashset.StringSet, error)
		GetClaims(userID string, profile string) (models.IClaims, error)
	}
)
