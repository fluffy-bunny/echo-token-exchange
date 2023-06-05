package probe

import "context"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/$GOPACKAGE IProbe

type (
	// IProbe ...
	IProbe interface {
		GetName() string
		Probe(ctx context.Context) error
	}
)
