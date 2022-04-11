package claimsprovider

import (
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IClaimsProvider"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IClaimsProvider

type (
	// IClaimsProvider ...
	IClaimsProvider interface {
		GetProfiles(userID string) ([]string, error)
		GetClaims(userID string, profile string) (contracts_tokenhandlers.Claims, error)
	}
)
