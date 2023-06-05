package removetokens

//go:generate mockgen -package=$GOPACKAGE -destination=../../../../mocks/background/tasks/$GOPACKAGE/mock_$GOFILE   echo-starter/internal/contracts/background/tasks/$GOPACKAGE IRemoveTokensSingletonTask
import (
	"context"
	contracts_background_tasks "echo-starter/internal/contracts/background/tasks"

	"github.com/hibiken/asynq"
)

const (
	TypeRemoveTokenByClientID           = "token:remove-by-client-id"
	TypeRemoveTokenBySubject            = "token:remove-by-subject"
	TypeRemoveTokenByClientIDAndSubject = "token:remove-by-client-id-and-subject"
)

type (
	TokenRemoveByClientID struct {
		ClientID string `json:"client_id"`
	}
	TokenRemoveBySubject struct {
		Subject string `json:"subject"`
	}
	TokenRemoveByClientIDAndSubject struct {
		TokenRemoveByClientID
		TokenRemoveBySubject
	}
	IRemoveTokensSingletonTask interface {
		contracts_background_tasks.ISingletonTask
		EnqueTaskTokenRemoveByClientID(ctx context.Context, task *TokenRemoveByClientID, opts ...asynq.Option) (*asynq.TaskInfo, error)
		EnqueTaskTypeRemoveTokenBySubject(ctx context.Context, task *TokenRemoveBySubject, opts ...asynq.Option) (*asynq.TaskInfo, error)
		EnqueTaskTokenRemoveByClientIDAndSubject(ctx context.Context, task *TokenRemoveByClientIDAndSubject, opts ...asynq.Option) (*asynq.TaskInfo, error)
	}
)
