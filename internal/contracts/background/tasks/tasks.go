package tasks

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ISingletonTask"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../../mocks/background/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/background/$GOPACKAGE ISingletonTask

import (
	"context"

	"github.com/hibiken/asynq"
)

type (
	ISingletonTask interface {
		ProcessTask(ctx context.Context, t *asynq.Task) error
		EnqueTask(payload interface{}) (*asynq.Task, error)
	}
)
