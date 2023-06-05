package rejson

// based off of
// https://github.com/go-oauth2/redis/blob/master/redis.go
import (
	"context"
	contracts_config "echo-starter/internal/contracts/config"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	"echo-starter/internal/models"
	"errors"
	"fmt"
	"sync"
	"time"

	di "github.com/dozm/di"
	core_utils "github.com/fluffy-bunny/fluffycore/utils"
	"github.com/fluffy-bunny/go-redis-search/ftsearch"
	"github.com/fluffy-bunny/rejonson/v8"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	jsonMarshal   = jsoniter.Marshal
	jsonUnmarshal = jsoniter.Unmarshal
)

type (
	service struct {
		Config         *contracts_config.Config `inject:"config"`
		lock           *sync.RWMutex
		opts           *redis.Options
		cli            *redis.Client
		ns             string
		rejonsonClient *rejonson.Client
		ftSearch       *ftsearch.Client
	}
	validated struct {
		scopes []string
	}
	entityTokenInfo struct {
		models.TokenInfo
		Key string `json:"key" redis:",key"` // the redis:",key" is required to indicate which field is the ULID key
		Ver int64  `json:"ver" redis:",ver"` // the redis:",ver" is required to do optimistic locking to prevent lost update
	}
)

var stemService *service = new(service)

func (s *service) Ctor(config *contracts_config.Config) (*service, error) {
	obj := &service{}
	obj.lock = &sync.RWMutex{}

	redisOptions := &redis.Options{
		Addr:     config.RedisOptions.Addr,
		Network:  config.RedisOptions.Network,
		Password: config.RedisOptions.Password,
		Username: config.RedisOptions.Username,
	}
	obj.cli = redis.NewClient(redisOptions)
	if len(config.RedisOptions.Namespace) > 0 {
		obj.ns = config.RedisOptions.Namespace[0]
	}

	obj.rejonsonClient = rejonson.ExtendClient(obj.cli)

	obj.ftSearch = ftsearch.NewClient(obj.cli)
	return obj, nil
}
func (s *service) Close() {
	s.cli.Close()

}
func init() {
	var _ contracts_stores_tokenstore.ITokenStore = (*service)(nil)
}

// AddSingletonITokenStore registers the *service.
func AddSingletonITokenStore(builder di.ContainerBuilder) {
	di.AddSingleton[contracts_stores_tokenstore.ITokenStore](builder, stemService.Ctor)
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

func (s *service) checkError(result redis.Cmder) (bool, error) {
	if err := result.Err(); err != nil {
		if err == redis.Nil {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
func (s *service) StoreToken(ctx context.Context, handle string, info *models.TokenInfo) (h string, err error) {
	log := zerolog.Ctx(ctx).With().Logger()
	//=================== PANIC RECOVERY ======================
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			log.Error().Err(err).Send()
		}
	}()
	//=================== PANIC RECOVERY ======================

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	if core_utils.IsEmptyOrNil(handle) {
		return "", errors.New("handle is empty")
	}

	basicID := handle
	expirationDuration := info.Metadata.Expiration.Sub(time.Now())
	aexp := expirationDuration
	rexp := aexp

	handleKey := s.wrapperKey(basicID)

	pipeline := s.rejonsonClient.Pipeline()
	json, err := jsonMarshal(info)
	pipeline.JsonSet(ctx, handleKey, ".", string(json))
	pipeline.Expire(ctx, handleKey, rexp)
	_, err = pipeline.Exec(ctx)
	if err != nil {
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

func (s *service) GetToken(ctx context.Context, handle string) (ti *models.TokenInfo, err error) {
	log := zerolog.Ctx(ctx).With().Logger()

	//=================== PANIC RECOVERY ======================
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			log.Error().Err(err).Send()
		}
	}()
	//=================== PANIC RECOVERY ======================

	if core_utils.IsEmptyOrNil(handle) {
		return nil, errors.New("handle is empty")
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.RLock()
	defer s.lock.RUnlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	handleKey := s.wrapperKey(handle)
	result := s.rejonsonClient.JsonGet(context.Background(), handleKey)
	if err != nil {
		// handle error
	}
	return s.parseToken(result)
}

// UpdateToken is like StoreToken except it only lets you do a StoreToken if the token already exists.
// no deep update logic here.
func (s *service) UpdateToken(ctx context.Context, handle string, info *models.TokenInfo) (err error) {
	log := zerolog.Ctx(ctx).With().Logger()

	//=================== PANIC RECOVERY ======================
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			log.Error().Err(err).Send()
		}
	}()
	//=================== PANIC RECOVERY ======================

	if core_utils.IsEmptyOrNil(handle) {
		return errors.New("handle is empty")
	}
	actual, err := s.GetToken(ctx, handle)
	if err != nil {
		return errors.New("not found")
	}
	if actual == nil {
		return errors.New("not found")
	}
	_, err = s.StoreToken(ctx, handle, info)
	return err

}
func (s *service) RemoveToken(ctx context.Context, handle string) (err error) {
	log := zerolog.Ctx(ctx).With().Logger()

	//=================== PANIC RECOVERY ======================
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			log.Error().Err(err).Send()
		}
	}()
	//=================== PANIC RECOVERY ======================

	if core_utils.IsEmptyOrNil(handle) {
		return errors.New("handle is empty")
	}
	// is it even here?
	info, err := s.GetToken(ctx, handle)
	if err != nil {
		return err
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
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

func (s *service) RemoveTokenByClientID(ctx context.Context, clientID string) (err error) {
	log := zerolog.Ctx(ctx).With().Logger()
	//=================== PANIC RECOVERY ======================
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			log.Error().Err(err).Send()
		}
	}()
	//=================== PANIC RECOVERY ======================

	if core_utils.IsEmptyOrNil(clientID) {
		return errors.New("client_id is empty")
	}

	var max int64 = 1000

	var totalProcessed int64 = 0
	for {
		queryString := fmt.Sprintf("@client_id:%s", clientID)
		query := ftsearch.NewQuery().
			WithIndex("echoTokenStoreIdx").
			WithQueryString(queryString).
			WithLimit(0, max)

		qResult, err := s.ftSearch.Search(ctx, query)
		if err != nil {
			return err
		}
		if qResult.Count == 0 {
			break
		}

		if qResult != nil {
			pipeline := s.rejonsonClient.Pipeline()
			for key := range qResult.Data {
				pipeline.Del(ctx, key)
			}
			_, err = pipeline.Exec(ctx)
		}
		totalProcessed += int64(len(qResult.Data))

	}

	return nil
}
func (s *service) RemoveTokenBySubject(ctx context.Context, subject string) (err error) {
	//=================== PANIC RECOVERY ======================
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			log.Error().Err(err).Send()
		}
	}()
	//=================== PANIC RECOVERY ======================

	if core_utils.IsEmptyOrNil(subject) {
		return errors.New("subject is empty")
	}

	var max int64 = 1000

	var totalProcessed int64 = 0
	for {
		queryString := fmt.Sprintf("@subject:%s", subject)
		query := ftsearch.NewQuery().
			WithIndex("echoTokenStoreIdx").
			WithQueryString(queryString).
			WithLimit(0, max)

		qResult, err := s.ftSearch.Search(ctx, query)
		if err != nil {
			return err
		}
		if qResult.Count == 0 {
			break
		}

		if qResult != nil {
			pipeline := s.rejonsonClient.Pipeline()
			for key := range qResult.Data {
				pipeline.Del(ctx, key)
			}
			_, err = pipeline.Exec(ctx)
		}
		totalProcessed += int64(len(qResult.Data))

	}

	return nil
}

func (s *service) RemoveTokenByClientIdAndSubject(ctx context.Context, clientID string, subject string) (err error) {
	//=================== PANIC RECOVERY ======================
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			log.Error().Err(err).Send()
		}
	}()
	//=================== PANIC RECOVERY ======================

	if core_utils.IsEmptyOrNil(clientID) {
		return errors.New("client_id is empty")
	}
	if core_utils.IsEmptyOrNil(subject) {
		return errors.New("subject is empty")
	}

	var max int64 = 1000

	var totalProcessed int64 = 0
	for {
		queryString := fmt.Sprintf("@client_id:%s,@subject:%s", clientID, subject)
		query := ftsearch.NewQuery().
			WithIndex("echoTokenStoreIdx").
			WithQueryString(queryString).
			WithLimit(0, max)

		qResult, err := s.ftSearch.Search(ctx, query)
		if err != nil {
			return err
		}
		if qResult.Count == 0 {
			break
		}

		if qResult != nil {
			pipeline := s.rejonsonClient.Pipeline()
			for key := range qResult.Data {
				pipeline.Del(ctx, key)
			}
			_, err = pipeline.Exec(ctx)
		}
		totalProcessed += int64(len(qResult.Data))

	}

	return nil
}
