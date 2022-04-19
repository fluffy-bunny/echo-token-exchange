package tasks

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ISingletonTask,ITaskEngine,ITaskClient"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../../mocks/background/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/background/$GOPACKAGE ISingletonTask,ITaskEngine,ITaskClient

import (
	"context"

	"github.com/hibiken/asynq"
)

type (
	ITaskEngine interface {
		Start() error
		Stop() error
	}
	ISingletonTask interface {
		GetPatterns() []string
		ProcessTask(ctx context.Context, t *asynq.Task) error
		EnqueTask(payload interface{}, opts ...asynq.Option) (*asynq.TaskInfo, error)
	}
	ITaskClient interface {
		EnqueTask(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
	}
)
