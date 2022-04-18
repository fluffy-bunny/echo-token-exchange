package taskegine

import (
	"reflect"

	contracts_background_tasks "echo-starter/internal/contracts/background/tasks"

	contracts_config "echo-starter/internal/contracts/config"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"

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
	service struct {
		Logger     contracts_logger.ILogger                    `inject:""`
		Config     *contracts_config.Config                    `inject:""`
		TokenStore contracts_stores_tokenstore.ITokenStore     `inject:""`
		Handlers   []contracts_background_tasks.ISingletonTask `inject:""`
		srv        *asynq.Server
		mux        *asynq.ServeMux
		future     async.Future
	}
)

func assertImplementation() {
	var _ contracts_background_tasks.ITaskEngine = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonITaskEngine registers the *service as a singleton.
func AddSingletonITaskEngine(builder *di.Builder) {
	contracts_background_tasks.AddSingletonITaskEngine(builder, reflectType)
}
func (s *service) Ctor() {
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
				"critical": 6,
				"default":  3,
				"low":      1,
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
	// add all the handlers
	for _, handler := range s.Handlers {
		for _, pattern := range handler.GetPatterns() {
			s.mux.Handle(pattern, handler)
		}
	}
	future := grpcdotnetgoasync.ExecuteWithPromiseAsync(func(promise async.Promise) {
		var err error
		defer func() {
			promise.Success(&grpcdotnetgoasync.AsyncResponse{
				Message: "End Serve - echo Server",
				Error:   err,
			})
		}()
		err = s.srv.Run(s.mux)
		if err != nil {
			s.Logger.Fatal().Err(err).Msg("Failed to start asynq server")
		}
	})
	s.future = future
	return nil
}

func (s *service) Stop() error {
	s.srv.Stop()
	promise, _ := s.future.Get()
	response := promise.(grpcdotnetgoasync.AsyncResponse)
	s.Logger.Info().Msg(response.Message)
	return nil
}
