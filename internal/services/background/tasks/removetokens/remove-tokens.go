package removetokens

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	contracts_background_tasks "echo-starter/internal/contracts/background/tasks"
	contracts_background_tasks_removetokens "echo-starter/internal/contracts/background/tasks/removetokens"

	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/hibiken/asynq"
)

const (
	TypeRemoveTokenByClientID = "token:remove-by-client-id"
)

type (
	service struct {
		Logger     contracts_logger.ILogger                `inject:""`
		TokenStore contracts_stores_tokenstore.ITokenStore `inject:""`
	}
)

func assertImplementation() {
	var _ contracts_background_tasks.ISingletonTask = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonISingletonTask registers the *service as a singleton.
func AddSingletonISingletonTask(builder *di.Builder) {
	contracts_background_tasks.AddSingletonISingletonTask(builder, reflectType)
}
func (s *service) processRemoveTokenByClientID(ctx context.Context, t *asynq.Task) error {

	var p contracts_background_tasks_removetokens.TokenRemoveByClientID
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		s.Logger.Error().Err(err).Msg("failed to unmarshal task payload")
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	err := s.TokenStore.RemoveTokenByClientID(ctx, p.ClientID)
	if err != nil {
		s.Logger.Error().Err(err).Msg("failed to remove token by client id")
		return fmt.Errorf("failed to remove token by client id: %v", err)
	}
	return nil
}
func (s *service) processRemoveTokenBySubject(ctx context.Context, t *asynq.Task) error {

	var p contracts_background_tasks_removetokens.TokenRemoveBySubject
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		s.Logger.Error().Err(err).Msg("failed to unmarshal task payload")
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	err := s.TokenStore.RemoveTokenBySubject(ctx, p.Subject)
	if err != nil {
		s.Logger.Error().Err(err).Msg("failed to remove token by subject")
		return fmt.Errorf("failed to remove token by subject: %v", err)
	}
	return nil
}
func (s *service) processRemoveTokenByClientIDAndSubject(ctx context.Context, t *asynq.Task) error {

	var p contracts_background_tasks_removetokens.TokenRemoveByClientIDAndSubject
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		s.Logger.Error().Err(err).Msg("failed to unmarshal task payload")
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	err := s.TokenStore.RemoveTokenByClientIdAndSubject(ctx, p.ClientID, p.Subject)
	if err != nil {
		s.Logger.Error().Err(err).Msg("failed to remove token by client id and subject")
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
func (s *service) EnqueTask(payload interface{}) (*asynq.Task, error) {
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
	return asynq.NewTask(name, payloadJson), nil
}
