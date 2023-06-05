package taskegine

import (
	"context"
	"reflect"

	contracts_background_tasks "echo-starter/internal/contracts/background/tasks"
	contracts_config "echo-starter/internal/contracts/config"

	di "github.com/dozm/di"
	asynq "github.com/hibiken/asynq"
	"github.com/rs/zerolog"
)

type (
	service struct {
		Config *contracts_config.Config `inject:""`
		client *asynq.Client
	}
)

func init() {
	var _ contracts_background_tasks.ITaskClient = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

func ctor(config *contracts_config.Config) (contracts_background_tasks.ITaskClient, asynq.ErrorHandler) {
	return &service{
		Config: config,
	}, nil
}

// AddSingletonITaskClient registers the *service as a singleton.
func AddSingletonITaskClient(builder di.ContainerBuilder) {
	di.AddSingleton[contracts_background_tasks.ITaskClient](builder,
		ctor)
}
func (s *service) Close() {
	s.client.Close()
}
func (s *service) Ctor() {
	s.client = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     s.Config.RedisOptions.Addr,
		Network:  s.Config.RedisOptions.Network,
		Password: s.Config.RedisOptions.Password,
		Username: s.Config.RedisOptions.Username})

}

func (s *service) EnqueTask(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	log := zerolog.Ctx(ctx).With().Logger()
	info, err := s.client.Enqueue(task, opts...)
	if err != nil {
		log.Error().Err(err).Msg("EnqueTask")
	}
	return info, err
}
