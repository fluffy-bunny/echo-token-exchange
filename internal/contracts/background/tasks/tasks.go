package tasks

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ISingletonTask,ITaskClient,ITaskEngineFactory"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/background/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/background/$GOPACKAGE ISingletonTask,ITaskClient,ITaskEngineFactory

import (
	"context"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	"github.com/hibiken/asynq"
)

const (
	TaskQueueTokenExchangeCritical = "tokenexchange:critical"
	TaskQueueTokenExchangeNormal   = "tokenexchange:normal"
	TaskQueueTokenExchangeLow      = "tokenexchange:low"
)

type (
	TaskEngineConfig struct {
		RedisClientOpt asynq.RedisClientOpt
		Config         asynq.Config
		Patterns       *core_hashset.StringSet
	}
	// ITaskEngineFactory
	ITaskEngineFactory interface {
		Start() error
		Stop() error
	}

	ISingletonTask interface {
		GetPatterns() *core_hashset.StringSet
		ProcessTask(ctx context.Context, t *asynq.Task) error
		EnqueTask(payload interface{}, opts ...asynq.Option) (*asynq.TaskInfo, error)
	}
	ITaskClient interface {
		EnqueTask(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
	}
)
