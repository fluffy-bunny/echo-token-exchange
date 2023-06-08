package inmemory

import (
	"context"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	"echo-starter/internal/models"
	"errors"
	"reflect"
	"sync"
	"time"

	di "github.com/dozm/di"
	core_utils "github.com/fluffy-bunny/fluffycore/utils"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

type (
	service struct {
		lock   *sync.RWMutex
		tokens map[string]*models.TokenInfo
	}
	validated struct {
		scopes []string
	}
)

var stemService *service = new(service)

func (s *service) Ctor() (*service, error) {
	obj := &service{
		lock:   &sync.RWMutex{},
		tokens: make(map[string]*models.TokenInfo),
	}
	return obj, nil
}
func init() {
	var _ contracts_stores_tokenstore.ITokenStore = (*service)(nil)
	var _ contracts_stores_tokenstore.IInternalTokenStore = (*service)(nil)

}

// AddSingletonITokenStore registers the *service.
func AddSingletonITokenStore(builder di.ContainerBuilder) {
	di.AddSingleton[*service](
		builder,
		stemService.Ctor,
		reflect.TypeOf((*contracts_stores_tokenstore.ITokenStore)(nil)),
		reflect.TypeOf((*contracts_stores_tokenstore.IInternalTokenStore)(nil)))

}

func (s *service) StoreToken(ctx context.Context, handle string, info *models.TokenInfo) (string, error) {

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	if core_utils.IsEmptyOrNil(handle) {
		return "", errors.New("handle is empty")
	}

	s.tokens[handle] = info
	return handle, nil
}
func (s *service) GetToken(ctx context.Context, handle string) (*models.TokenInfo, error) {

	if core_utils.IsEmptyOrNil(handle) {
		return nil, errors.New("handle is empty")
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.RLock()
	defer s.lock.RUnlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	h, ok := s.tokens[handle]
	if !ok {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return h, nil

}
func (s *service) UpdateToken(ctx context.Context, handle string, info *models.TokenInfo) error {

	if core_utils.IsEmptyOrNil(handle) {
		return errors.New("handle is empty")
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	_, ok := s.tokens[handle]
	if !ok {
		return nil
	}

	s.tokens[handle] = info
	return nil
}
func (s *service) RemoveToken(ctx context.Context, handle string) error {

	if core_utils.IsEmptyOrNil(handle) {
		return errors.New("handle is empty")
	}
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	_, ok := s.tokens[handle]
	if !ok {
		return nil
	}
	delete(s.tokens, handle)
	return nil
}
func (s *service) RemoveTokenByClientID(ctx context.Context, clientID string) error {

	if core_utils.IsEmptyOrNil(clientID) {
		return errors.New("client_id is empty")
	}
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	for k, v := range s.tokens {
		if v.Metadata.ClientID == clientID {
			delete(s.tokens, k)
		}
	}
	return nil
}
func (s *service) RemoveTokenBySubject(ctx context.Context, subject string) error {

	if core_utils.IsEmptyOrNil(subject) {
		return errors.New("subject is empty")
	}
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--

	for k, v := range s.tokens {
		if v.Metadata.Subject == subject {
			delete(s.tokens, k)
		}
	}
	return nil
}
func (s *service) RemoveTokenByClientIdAndSubject(ctx context.Context, clientID string, subject string) error {

	if core_utils.IsEmptyOrNil(clientID) {
		return errors.New("client_id is empty")
	}
	if core_utils.IsEmptyOrNil(subject) {
		return errors.New("subject is empty")
	}
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	for k, v := range s.tokens {
		if v.Metadata.ClientID == clientID && v.Metadata.Subject == subject {
			delete(s.tokens, k)
		}
	}
	return nil
}
func (s *service) RemoveExpired(ctx context.Context) error {

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.RLock()
	defer s.lock.RUnlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	now := time.Now()
	var handles []string
	for k, v := range s.tokens {
		if now.After(v.Metadata.Expiration) {
			handles = append(handles, k)
		}
	}
	if len(handles) > 0 {
		go func() {
			//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
			s.lock.Lock()
			defer s.lock.Unlock()
			//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
			for _, v := range handles {
				delete(s.tokens, v)
			}
		}()
	}
	return nil
}
