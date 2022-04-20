package taskenginefactory

import (
	"reflect"

	contracts_background_tasks "echo-starter/internal/contracts/background/tasks"

	contracts_background_tasks_removetokens "echo-starter/internal/contracts/background/tasks/removetokens"
	contracts_config "echo-starter/internal/contracts/config"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"

	grpcdotnetgoasync "github.com/fluffy-bunny/grpcdotnetgo/pkg/async"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/hibiken/asynq"
	"github.com/reugn/async"
)

const (
	TypeRemoveTokenByClientID = "token:remove-by-client-id"
)

type (
	serverMuxContainer struct {
		config contracts_background_tasks.TaskEngineConfig
		mux    *asynq.ServeMux
		srv    *asynq.Server
		future async.Future
	}
	service struct {
		Logger              contracts_logger.ILogger                    `inject:""`
		Config              *contracts_config.Config                    `inject:""`
		Handlers            []contracts_background_tasks.ISingletonTask `inject:""`
		srv                 *asynq.Server
		mux                 *asynq.ServeMux
		future              async.Future
		taskEngineConfigs   []contracts_background_tasks.TaskEngineConfig
		serverMuxContainers []*serverMuxContainer
	}
)

func assertImplementation() {
	var _ contracts_background_tasks.ITaskEngineFactory = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonITaskEngine registers the *service as a singleton.
func AddSingletonITaskEngine(builder *di.Builder) {
	contracts_background_tasks.AddSingletonITaskEngine(builder, reflectType)
}
func (s *service) Ctor() {
	s.taskEngineConfigs = append(s.taskEngineConfigs, contracts_background_tasks.TaskEngineConfig{
		RedisClientOpt: asynq.RedisClientOpt{Addr: s.Config.RedisOptionsReferenceTokenStore.Addr,
			Network:  s.Config.RedisOptionsReferenceTokenStore.Network,
			Password: s.Config.RedisOptionsReferenceTokenStore.Password,
			Username: s.Config.RedisOptionsReferenceTokenStore.Username,
		},
		Config: asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				contracts_background_tasks.TaskQueueTokenExchangeCritical: 6,
				contracts_background_tasks.TaskQueueTokenExchangeNormal:   3,
				contracts_background_tasks.TaskQueueTokenExchangeLow:      1,
			},
		},
		Patterns: core_hashset.NewStringSet(
			contracts_background_tasks_removetokens.TypeRemoveTokenByClientID,
			contracts_background_tasks_removetokens.TypeRemoveTokenBySubject,
			contracts_background_tasks_removetokens.TypeRemoveTokenByClientIDAndSubject),
	})
	for _, config := range s.taskEngineConfigs {
		srv := asynq.NewServer(config.RedisClientOpt, config.Config)
		mux := asynq.NewServeMux()
		s.serverMuxContainers = append(s.serverMuxContainers, &serverMuxContainer{
			config: config,
			mux:    mux,
			srv:    srv,
		})
	}
	for _, container := range s.serverMuxContainers {
		for _, handler := range s.Handlers {
			for _, pattern := range handler.GetPatterns().Values() {
				if container.config.Patterns.Contains(pattern) {
					container.mux.Handle(pattern, handler)
				}
			}
		}
	}

	s.srv = asynq.NewServer(
		asynq.RedisClientOpt{Addr: s.Config.RedisOptionsReferenceTokenStore.Addr,
			Network:  s.Config.RedisOptionsReferenceTokenStore.Network,
			Password: s.Config.RedisOptionsReferenceTokenStore.Password,
			Username: s.Config.RedisOptionsReferenceTokenStore.Username},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				contracts_background_tasks.TaskQueueTokenExchangeCritical: 6,
				contracts_background_tasks.TaskQueueTokenExchangeNormal:   3,
				contracts_background_tasks.TaskQueueTokenExchangeLow:      1,
			},
			// See the godoc for other configuration options
		},
	)
	s.mux = asynq.NewServeMux()
}

func (s *service) Start() error {

	if s.future != nil {
		panic("task engine already started")
	}

	for _, container := range s.serverMuxContainers {
		container.future = grpcdotnetgoasync.ExecuteWithPromiseAsync(func(promise async.Promise) {
			var err error
			defer func() {
				promise.Success(&grpcdotnetgoasync.AsyncResponse{
					Message: "End Serve - echo Server",
					Error:   err,
				})
			}()
			err = container.srv.Run(container.mux)
			if err != nil {
				s.Logger.Fatal().Err(err).Msg("Failed to start asynq server")
			}
		})
	}
	return nil
}

func (s *service) Stop() error {
	// tell all to close
	for _, container := range s.serverMuxContainers {
		container.srv.Stop()
	}

	// wait for all to return the promise
	for _, container := range s.serverMuxContainers {
		promise, _ := container.future.Get()
		response := promise.(grpcdotnetgoasync.AsyncResponse)
		s.Logger.Info().Msg(response.Message)
	}
	return nil
}
