package claimsprovider

import (
	"echo-starter/internal/models"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IClaimsProvider"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IClaimsProvider

type (
	// IClaimsProvider ...
	IClaimsProvider interface {
		GetProfiles(userID string) (*core_hashset.StringSet, error)
		GetClaims(userID string, profile string) (models.IClaims, error)
	}
)
