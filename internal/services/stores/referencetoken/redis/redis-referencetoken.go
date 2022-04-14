package redis

// based off of
// https://github.com/go-oauth2/redis/blob/master/redis.go
import (
	"context"
	contracts_config "echo-starter/internal/contracts/config"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/xid"
)

var (
	jsonMarshal   = jsoniter.Marshal
	jsonUnmarshal = jsoniter.Unmarshal
)

type (
	service struct {
		Config *contracts_config.Config `inject:"config"`
		lock   *sync.RWMutex
		tokens map[string]*contracts_stores_tokenstore.ReferenceTokenInfo
		opts   *redis.Options
		cli    *redis.Client
		ns     string
	}
	validated struct {
		scopes []string
	}
)

func (s *service) Ctor() {
	s.lock = &sync.RWMutex{}
	s.cli = redis.NewClient(&redis.Options{
		Addr:     s.Config.RedisOptionsReferenceTokenStore.Addr,
		Network:  s.Config.RedisOptionsReferenceTokenStore.Network,
		Password: s.Config.RedisOptionsReferenceTokenStore.Password,
		Username: s.Config.RedisOptionsReferenceTokenStore.Username,
	})
	if len(s.Config.RedisOptionsReferenceTokenStore.Namespace) > 0 {
		s.ns = s.Config.RedisOptionsReferenceTokenStore.Namespace[0]
	}
	s.tokens = make(map[string]*contracts_stores_tokenstore.ReferenceTokenInfo)
}
func (s *service) Close() {
	s.cli.Close()
}
func assertImplementation() {
	var _ contracts_stores_tokenstore.IReferenceTokenStore = (*service)(nil)
	var _ contracts_stores_tokenstore.IInternalReferenceTokenStore = (*service)(nil)

}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIReferenceTokenStore registers the *service.
func AddSingletonIReferenceTokenStore(builder *di.Builder) {
	contracts_stores_tokenstore.AddSingletonIReferenceTokenStore(builder, reflectType,
		contracts_stores_tokenstore.ReflectTypeIInternalReferenceTokenStore)
}
func (s *service) wrapClientIDKey(clientID string) string {
	return fmt.Sprintf("%s:client_id:%s", s.ns, clientID)
}
func (s *service) wrapSubjectKey(subject string) string {
	return fmt.Sprintf("%s:subject:%s", s.ns, subject)
}

func (s *service) wrapClientSubjectKey(clientID string, subject string) string {
	return fmt.Sprintf("%s:client_id:%s:subject:%s", s.ns, clientID, subject)
}
func (s *service) wrapperKey(key string) string {
	return fmt.Sprintf("%s%s", s.ns, key)
}
func (s *service) checkError(result redis.Cmder) (bool, error) {
	if err := result.Err(); err != nil {
		if err == redis.Nil {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
func (s *service) StoreReferenceToken(ctx context.Context, info *contracts_stores_tokenstore.ReferenceTokenInfo) (handle string, err error) {

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	handle = xid.New().String()
	jv, err := jsonMarshal(info)
	if err != nil {
		return "", err
	}
	pipe := s.cli.TxPipeline()

	basicID := handle
	expirationDuration := info.Expiration.Sub(time.Now())
	aexp := expirationDuration
	rexp := aexp

	handleKey := s.wrapperKey(basicID)

	// we need to know every handle that was created on behalf of a client
	pipe.SAdd(ctx, s.wrapClientIDKey(info.ClientID), handleKey)
	if !core_utils.IsEmptyOrNil(info.Subject) {
		// we need to know every handle that was created on behalf of a subject
		pipe.SAdd(ctx, s.wrapSubjectKey(info.Subject), handleKey)
		clientSubjectKey := s.wrapClientSubjectKey(info.ClientID, info.Subject)
		pipe.Set(ctx, clientSubjectKey, handleKey, rexp)
	}
	// finally recored the token itself
	pipe.Set(ctx, handleKey, jv, rexp)

	if _, err := pipe.Exec(ctx); err != nil {
		return "", err
	}

	return handle, nil
}
func (s *service) parseToken(result *redis.StringCmd) (*contracts_stores_tokenstore.ReferenceTokenInfo, error) {
	if ok, err := s.checkError(result); err != nil {
		return nil, err
	} else if ok {
		return nil, nil
	}

	buf, err := result.Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	token := contracts_stores_tokenstore.ReferenceTokenInfo{}
	if err := jsonUnmarshal(buf, &token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *service) GetReferenceToken(ctx context.Context, handle string) (*contracts_stores_tokenstore.ReferenceTokenInfo, error) {

	if core_utils.IsEmptyOrNil(handle) {
		return nil, errors.New("handle is empty")
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.RLock()
	defer s.lock.RUnlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	result := s.cli.Get(ctx, s.wrapperKey(handle))
	return s.parseToken(result)
}
func (s *service) UpdateReferenceToken(ctx context.Context, handle string, info *contracts_stores_tokenstore.ReferenceTokenInfo) error {

	if core_utils.IsEmptyOrNil(handle) {
		return errors.New("handle is empty")
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	result := s.cli.Get(ctx, s.wrapperKey(handle))
	_, err := s.parseToken(result)
	if err != nil {
		return errors.New("not found")
	}
	jv, err := jsonMarshal(info)
	if err != nil {
		return err
	}
	pipe := s.cli.TxPipeline()

	basicID := handle
	expirationDuration := info.Expiration.Sub(time.Now())
	aexp := expirationDuration
	rexp := aexp

	// finally recored the token itself
	// we only need to update the main key
	pipe.Set(ctx, s.wrapperKey(basicID), jv, rexp)

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}
func (s *service) RemoveReferenceToken(ctx context.Context, handle string) error {

	if core_utils.IsEmptyOrNil(handle) {
		return errors.New("handle is empty")
	}
	// is it even here?
	info, err := s.GetReferenceToken(ctx, handle)
	if err != nil {
		return err
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--

	pipe := s.cli.TxPipeline()

	// we need to know every handle that was created on behalf of a client
	pipe.SRem(ctx, s.wrapClientIDKey(info.ClientID), s.wrapperKey(handle))
	if !core_utils.IsEmptyOrNil(info.Subject) {
		// we need to know every handle that was created on behalf of a subject
		pipe.SRem(ctx, s.wrapSubjectKey(info.Subject), s.wrapperKey(handle))
		clientSubjectKey := s.wrapClientSubjectKey(info.ClientID, info.Subject)
		pipe.Del(ctx, clientSubjectKey)
	}
	// finally delete the token itself
	pipe.Del(ctx, s.wrapperKey(handle))

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}
func (s *service) removeOneSetBlock(ctx context.Context, setKey string) (more bool, err error) {
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--

	keys, _, err := s.cli.SScan(ctx, setKey, 0, "", 0).Result()
	if err != nil {
		return false, err
	}
	if len(keys) == 0 {
		return false, nil
	}
	pipe := s.cli.TxPipeline()
	for _, key := range keys {
		pipe.Del(ctx, key)
		pipe.SRem(ctx, setKey, key)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return false, err
	}
	return true, nil
}

func (s *service) RemoveReferenceTokenByClientID(ctx context.Context, clientID string) error {
	if core_utils.IsEmptyOrNil(clientID) {
		return errors.New("client_id is empty")
	}
	setKey := s.wrapClientIDKey(clientID)
	for {
		more, err := s.removeOneSetBlock(ctx, setKey)
		if err != nil {
			return err
		}
		if !more {
			break
		}
	}

	return nil
}
func (s *service) RemoveReferenceTokenBySubject(ctx context.Context, subject string) error {
	if core_utils.IsEmptyOrNil(subject) {
		return errors.New("subject is empty")
	}
	setKey := s.wrapClientIDKey(subject)
	for {
		more, err := s.removeOneSetBlock(ctx, setKey)
		if err != nil {
			return err
		}
		if !more {
			break
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
	clientSubjectKey := s.wrapClientSubjectKey(clientID, subject)
	handleKey := s.cli.Get(ctx, clientSubjectKey).String()
	if core_utils.IsEmptyOrNil(handleKey) {
		return errors.New("not found")
	}
	pipe := s.cli.TxPipeline()

	// we need to know every handle that was created on behalf of a client
	pipe.SRem(ctx, s.wrapClientIDKey(clientID), handleKey)
	pipe.SRem(ctx, s.wrapSubjectKey(subject), handleKey)
	pipe.Del(ctx, clientSubjectKey)
	// finally delete the token itself
	pipe.Del(ctx, handleKey)

	if _, err := pipe.Exec(ctx); err != nil {
		return err
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
