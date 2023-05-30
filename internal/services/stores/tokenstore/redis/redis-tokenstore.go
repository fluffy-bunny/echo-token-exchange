package redis

// based off of
// https://github.com/go-oauth2/redis/blob/master/redis.go
import (
	"context"
	contracts_config "echo-starter/internal/contracts/config"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	"echo-starter/internal/models"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/bsm/redislock"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
)

var (
	jsonMarshal   = jsoniter.Marshal
	jsonUnmarshal = jsoniter.Unmarshal
)

type (
	service struct {
		Config *contracts_config.Config `inject:"config"`
		lock   *sync.RWMutex
		opts   *redis.Options
		cli    *redis.Client
		ns     string
		locker *redislock.Client
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
	// Create a new lock client.
	s.locker = redislock.New(s.cli)
}
func (s *service) Close() {
	s.cli.Close()
}
func assertImplementation() {
	var _ contracts_stores_tokenstore.ITokenStore = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonITokenStore registers the *service.
func AddSingletonITokenStore(builder *di.Builder) {
	contracts_stores_tokenstore.AddSingletonITokenStore(builder, reflectType)
}
func (s *service) obtainLock(ctx context.Context, originalKey string, millSeconds int) (*redislock.Lock, error) {
	return s.locker.Obtain(ctx, s.wrapperLockKey(originalKey), time.Duration(millSeconds)*time.Millisecond, nil)
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
	return fmt.Sprintf("%s:%s", s.ns, key)
}
func (s *service) wrapperLockKey(key string) string {
	return fmt.Sprintf("%s:%s:lock", s.ns, key)
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
func (s *service) StoreToken(ctx context.Context, handle string, info *models.TokenInfo) (string, error) {

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	//s.lock.Lock()
	//defer s.lock.Unlock()
	lock, err := s.obtainLock(ctx, handle)
	if err != nil {
		return "", err
	}
	defer lock.Release(ctx)
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	if core_utils.IsEmptyOrNil(handle) {
		return "", errors.New("handle is empty")
	}

	jv, err := jsonMarshal(info)
	if err != nil {
		return "", err
	}
	pipe := s.cli.TxPipeline()

	basicID := handle
	expirationDuration := info.Metadata.Expiration.Sub(time.Now())
	aexp := expirationDuration
	rexp := aexp

	handleKey := s.wrapperKey(basicID)

	// we need to know every handle that was created on behalf of a client
	pipe.SAdd(ctx, s.wrapClientIDKey(info.Metadata.ClientID), handleKey)
	if !core_utils.IsEmptyOrNil(info.Metadata.Subject) {
		// we need to know every handle that was created on behalf of a subject
		pipe.SAdd(ctx, s.wrapSubjectKey(info.Metadata.Subject), handleKey)
		clientSubjectKey := s.wrapClientSubjectKey(info.Metadata.ClientID, info.Metadata.Subject)
		pipe.Set(ctx, clientSubjectKey, handleKey, rexp)
	}
	// finally recored the token itself
	pipe.Set(ctx, handleKey, jv, rexp)

	if _, err := pipe.Exec(ctx); err != nil {
		return "", err
	}

	return handle, nil
}
func (s *service) parseToken(result *redis.StringCmd) (*models.TokenInfo, error) {
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

	token := models.TokenInfo{}
	if err := jsonUnmarshal(buf, &token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *service) GetToken(ctx context.Context, handle string) (*models.TokenInfo, error) {

	if core_utils.IsEmptyOrNil(handle) {
		return nil, errors.New("handle is empty")
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	//s.lock.Lock()
	//defer s.lock.Unlock()
	lock, err := s.obtainLock(ctx, handle)
	if err != nil {
		return nil, err
	}
	defer lock.Release(ctx)
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	result := s.cli.Get(ctx, s.wrapperKey(handle))
	return s.parseToken(result)
}
func (s *service) UpdateToken(ctx context.Context, handle string, info *models.TokenInfo) error {

	if core_utils.IsEmptyOrNil(handle) {
		return errors.New("handle is empty")
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	//s.lock.Lock()
	//defer s.lock.Unlock()
	lock, err := s.obtainLock(ctx, handle)
	if err != nil {
		return err
	}
	defer lock.Release(ctx)
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	result := s.cli.Get(ctx, s.wrapperKey(handle))
	_, err = s.parseToken(result)
	if err != nil {
		return errors.New("not found")
	}
	jv, err := jsonMarshal(info)
	if err != nil {
		return err
	}
	pipe := s.cli.TxPipeline()

	basicID := handle
	expirationDuration := info.Metadata.Expiration.Sub(time.Now())
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
func (s *service) RemoveToken(ctx context.Context, handle string) error {

	if core_utils.IsEmptyOrNil(handle) {
		return errors.New("handle is empty")
	}
	// is it even here?
	info, err := s.GetToken(ctx, handle)
	if err != nil {
		return err
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	//s.lock.Lock()
	//defer s.lock.Unlock()
	lock, err := s.obtainLock(ctx, handle)
	if err != nil {
		return err
	}
	defer lock.Release(ctx)
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--

	pipe := s.cli.TxPipeline()

	// we need to know every handle that was created on behalf of a client
	pipe.SRem(ctx, s.wrapClientIDKey(info.Metadata.ClientID), s.wrapperKey(handle))
	if !core_utils.IsEmptyOrNil(info.Metadata.Subject) {
		// we need to know every handle that was created on behalf of a subject
		pipe.SRem(ctx, s.wrapSubjectKey(info.Metadata.Subject), s.wrapperKey(handle))
		clientSubjectKey := s.wrapClientSubjectKey(info.Metadata.ClientID, info.Metadata.Subject)
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

func (s *service) RemoveTokenByClientID(ctx context.Context, clientID string) error {
	if core_utils.IsEmptyOrNil(clientID) {
		return errors.New("client_id is empty")
	}
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	//s.lock.Lock()
	//defer s.lock.Unlock()
	lock, err := s.obtainLock(ctx, clientID)
	if err != nil {
		return err
	}
	defer lock.Release(ctx)
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
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
func (s *service) RemoveTokenBySubject(ctx context.Context, subject string) error {
	if core_utils.IsEmptyOrNil(subject) {
		return errors.New("subject is empty")
	}
	setKey := s.wrapSubjectKey(subject)
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
	clientSubjectKey := s.wrapClientSubjectKey(clientID, subject)
	handleKey, _ := s.cli.Get(ctx, clientSubjectKey).Result()
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
