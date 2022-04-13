package inmemory

import (
	"context"
	contracts_stores_referencetoken "echo-starter/internal/contracts/stores/referencetoken"
	"errors"
	"reflect"
	"sync"
	"time"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/xid"
)

type (
	service struct {
		lock   *sync.RWMutex
		tokens map[string]*contracts_stores_referencetoken.ReferenceTokenInfo
	}
	validated struct {
		scopes []string
	}
)

func (s *service) Ctor() {
	s.lock = &sync.RWMutex{}
	s.tokens = make(map[string]*contracts_stores_referencetoken.ReferenceTokenInfo)
}
func assertImplementation() {
	var _ contracts_stores_referencetoken.IReferenceTokenStore = (*service)(nil)
	var _ contracts_stores_referencetoken.IInternalReferenceTokenStore = (*service)(nil)

}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIReferenceTokenStore registers the *service.
func AddSingletonIReferenceTokenStore(builder *di.Builder) {
	contracts_stores_referencetoken.AddSingletonIReferenceTokenStore(builder, reflectType,
		contracts_stores_referencetoken.ReflectTypeIInternalReferenceTokenStore)
}
func (s *service) StoreReferenceToken(ctx context.Context, info *contracts_stores_referencetoken.ReferenceTokenInfo) (handle string, err error) {

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	handle = xid.New().String()
	s.tokens[handle] = info
	return handle, nil
}
func (s *service) GetReferenceToken(ctx context.Context, handle string) (*contracts_stores_referencetoken.ReferenceTokenInfo, error) {

	if core_utils.IsEmptyOrNil(handle) {
		return nil, errors.New("handle is empty")
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.RLock()
	defer s.lock.RUnlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	h, ok := s.tokens[handle]
	if !ok {
		return nil, errors.New("not found")
	}
	return h, nil

}
func (s *service) UpdateReferenceToken(ctx context.Context, handle string, info *contracts_stores_referencetoken.ReferenceTokenInfo) error {

	if core_utils.IsEmptyOrNil(handle) {
		return errors.New("handle is empty")
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	_, ok := s.tokens[handle]
	if !ok {
		return errors.New("not found")
	}

	s.tokens[handle] = info
	return nil
}
func (s *service) RemoveReferenceToken(ctx context.Context, handle string) error {

	if core_utils.IsEmptyOrNil(handle) {
		return errors.New("handle is empty")
	}
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	_, ok := s.tokens[handle]
	if !ok {
		return errors.New("not found")
	}
	delete(s.tokens, handle)
	return nil
}
func (s *service) RemoveReferenceTokenByClientID(ctx context.Context, clientID string) error {

	if core_utils.IsEmptyOrNil(clientID) {
		return errors.New("client_id is empty")
	}
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	for k, v := range s.tokens {
		if v.ClientID == clientID {
			delete(s.tokens, k)
		}
	}
	return nil
}
func (s *service) RemoveReferenceTokenBySubject(ctx context.Context, subject string) error {

	if core_utils.IsEmptyOrNil(subject) {
		return errors.New("subject is empty")
	}
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--

	for k, v := range s.tokens {
		if v.Subject == subject {
			delete(s.tokens, k)
		}
	}
	return nil
}
func (s *service) RemoveReferenceTokenByClientIdAndSubject(ctx context.Context, clientID string, subject string) error {

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
		if v.ClientID == clientID && v.Subject == subject {
			delete(s.tokens, k)
		}
	}
	return nil
}
func (s *service) RemoveExpired(ctx context.Context) {

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.RLock()
	defer s.lock.RUnlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	now := time.Now()
	var handles []string
	for k, v := range s.tokens {
		if now.After(v.Expiration) {
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

}
