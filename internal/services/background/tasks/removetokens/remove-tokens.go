package removetokens

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	contracts_background_tasks "echo-starter/internal/contracts/background/tasks"
	contracts_background_tasks_removetokens "echo-starter/internal/contracts/background/tasks/removetokens"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	"github.com/rs/zerolog"

	di "github.com/dozm/di"
	"github.com/hibiken/asynq"
)

const (
	TypeRemoveTokenByClientID = "token:remove-by-client-id"
)

type (
	service struct {
		TokenStore contracts_stores_tokenstore.ITokenStore `inject:""`
		TaskClient contracts_background_tasks.ITaskClient  `inject:""`
	}
)

func init() {
	var _ contracts_background_tasks.ISingletonTask = (*service)(nil)
	var _ contracts_background_tasks_removetokens.IRemoveTokensSingletonTask = (*service)(nil)

}

var reflectType = reflect.TypeOf((*service)(nil))

func ctor(tokenStore contracts_stores_tokenstore.ITokenStore, taskClient contracts_background_tasks.ITaskClient) (*service, error) {
	return &service{
		TokenStore: tokenStore,
		TaskClient: taskClient,
	}, nil
}

// AddSingletonISingletonTask registers the *service as a singleton.
func AddSingletonISingletonTask(builder di.ContainerBuilder) {

	di.AddSingleton[*service](builder, ctor,
		reflect.TypeOf((*contracts_background_tasks.ISingletonTask)(nil)),
		reflect.TypeOf((*contracts_background_tasks_removetokens.IRemoveTokensSingletonTask)(nil)),
	)
}
func (s *service) GetPatterns() *core_hashset.StringSet {
	return core_hashset.NewStringSet(
		contracts_background_tasks_removetokens.TypeRemoveTokenByClientID,
		contracts_background_tasks_removetokens.TypeRemoveTokenBySubject,
		contracts_background_tasks_removetokens.TypeRemoveTokenByClientIDAndSubject)

}
func (s *service) processRemoveTokenByClientID(ctx context.Context, t *asynq.Task) error {
	log := zerolog.Ctx(ctx).With().Logger()
	var p contracts_background_tasks_removetokens.TokenRemoveByClientID
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal task payload")
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	err := s.TokenStore.RemoveTokenByClientID(ctx, p.ClientID)
	if err != nil {
		log.Error().Err(err).Msg("failed to remove token by client id")
		return fmt.Errorf("failed to remove token by client id: %v", err)
	}
	return nil
}
func (s *service) processRemoveTokenBySubject(ctx context.Context, t *asynq.Task) error {
	log := zerolog.Ctx(ctx).With().Logger()

	var p contracts_background_tasks_removetokens.TokenRemoveBySubject
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal task payload")
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	err := s.TokenStore.RemoveTokenBySubject(ctx, p.Subject)
	if err != nil {
		log.Error().Err(err).Msg("failed to remove token by subject")
		return fmt.Errorf("failed to remove token by subject: %v", err)
	}
	return nil
}
func (s *service) processRemoveTokenByClientIDAndSubject(ctx context.Context, t *asynq.Task) error {
	log := zerolog.Ctx(ctx).With().Logger()

	var p contracts_background_tasks_removetokens.TokenRemoveByClientIDAndSubject
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal task payload")
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	err := s.TokenStore.RemoveTokenByClientIdAndSubject(ctx, p.ClientID, p.Subject)
	if err != nil {
		log.Error().Err(err).Msg("failed to remove token by client id and subject")
		return fmt.Errorf("failed to remove token by client id and subject: %v", err)
	}
	return nil
}
func (s *service) ProcessTask(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case contracts_background_tasks_removetokens.TypeRemoveTokenByClientID:
		return s.processRemoveTokenByClientID(ctx, t)
	case contracts_background_tasks_removetokens.TypeRemoveTokenBySubject:
		return s.processRemoveTokenBySubject(ctx, t)
	case contracts_background_tasks_removetokens.TypeRemoveTokenByClientIDAndSubject:
		return s.processRemoveTokenByClientIDAndSubject(ctx, t)
	default:
		return fmt.Errorf("unknown task type: %s", t.Type())
	}

}
func (s *service) EnqueTask(ctx context.Context, payload interface{}, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	var name string
	if removeByClient := payload.(*contracts_background_tasks_removetokens.TokenRemoveByClientID); removeByClient != nil {
		name = contracts_background_tasks_removetokens.TypeRemoveTokenByClientID
	} else if removeBySubject := payload.(*contracts_background_tasks_removetokens.TokenRemoveBySubject); removeBySubject != nil {
		name = contracts_background_tasks_removetokens.TypeRemoveTokenBySubject
	} else if removeByClientIdAndSubject := payload.(*contracts_background_tasks_removetokens.TokenRemoveByClientIDAndSubject); removeByClientIdAndSubject != nil {
		name = contracts_background_tasks_removetokens.TypeRemoveTokenByClientIDAndSubject
	} else {
		return nil, fmt.Errorf("invalid payload type: %v", reflect.TypeOf(payload))
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(name, payloadJson)
	return s.TaskClient.EnqueTask(ctx, task, opts...)
}
func (s *service) EnqueTaskTokenRemoveByClientID(ctx context.Context, task *contracts_background_tasks_removetokens.TokenRemoveByClientID, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return s.EnqueTask(ctx, task, opts...)
}
func (s *service) EnqueTaskTypeRemoveTokenBySubject(ctx context.Context, task *contracts_background_tasks_removetokens.TokenRemoveBySubject, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return s.EnqueTask(ctx, task, opts...)
}
func (s *service) EnqueTaskTokenRemoveByClientIDAndSubject(ctx context.Context, task *contracts_background_tasks_removetokens.TokenRemoveByClientIDAndSubject, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return s.EnqueTask(ctx, task, opts...)
}
