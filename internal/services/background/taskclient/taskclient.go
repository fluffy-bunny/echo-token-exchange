package taskegine

import (
	"reflect"

	contracts_background_tasks "echo-starter/internal/contracts/background/tasks"

	contracts_config "echo-starter/internal/contracts/config"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/hibiken/asynq"
)

type (
	service struct {
		Logger contracts_logger.ILogger `inject:""`
		Config *contracts_config.Config `inject:""`
		client *asynq.Client
	}
)

func assertImplementation() {
	var _ contracts_background_tasks.ITaskClient = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonITaskClient registers the *service as a singleton.
func AddSingletonITaskClient(builder *di.Builder) {
	contracts_background_tasks.AddSingletonITaskClient(builder, reflectType)
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

func (s *service) EnqueTask(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	info, err := s.client.Enqueue(task, opts...)
	if err != nil {
		s.Logger.Error().Err(err).Msg("EnqueTask")
	}
	return info, err
}
