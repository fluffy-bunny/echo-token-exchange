package jwt

import (
	"context"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	"echo-starter/internal/models"
	"encoding/json"
	"errors"
	"reflect"
	"sync"
	"time"

	di "github.com/dozm/di"
	fluffycore_contracts_common "github.com/fluffy-bunny/fluffycore/contracts/common"
	core_hashset "github.com/fluffy-bunny/fluffycore/gods/sets/hashset"
	core_utils "github.com/fluffy-bunny/fluffycore/utils"
	"github.com/golang-jwt/jwt"
	"github.com/rs/xid"
)

type (
	service struct {
		JWTTokenStore contracts_stores_tokenstore.IJwtTokenStore `inject:""`
		Now           fluffycore_contracts_common.TimeNow        `inject:""`

		lock   *sync.RWMutex
		tokens map[string]*models.TokenInfo
	}
	validated struct {
		scopes []string
	}
)

var stemService *service = new(service)

func (s *service) Ctor(
	now fluffycore_contracts_common.TimeNow,
	jwtTokenStore contracts_stores_tokenstore.IJwtTokenStore) (*service, error) {
	obj := &service{
		Now:           now,
		JWTTokenStore: jwtTokenStore,
		lock:          &sync.RWMutex{},
		tokens:        make(map[string]*models.TokenInfo),
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
	var buildClaimsMap = func(ctx context.Context, standardClaims *jwt.StandardClaims, extras models.IClaims) models.IClaims {
		audienceSet := core_hashset.NewStringSet()
		if !core_utils.IsEmptyOrNil(standardClaims.Audience) {
			audienceSet.Add(standardClaims.Audience)
		}
		if !core_utils.IsNil(extras) {
			extraAudInterface := extras.Get("aud")
			switch extraAudienceType := extraAudInterface.(type) {
			case string:
				audienceSet.Add(extraAudienceType)
			case []string:
				audienceSet.Add(extraAudienceType...)
			}
		}
		extras.Set("aud", audienceSet.Values())

		var standard map[string]interface{}
		standardJSON, _ := json.Marshal(standardClaims)
		json.Unmarshal(standardJSON, &standard)
		delete(standard, "aud")

		for k, v := range standard {
			extras.Set(k, v)
		}
		return extras
	}
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.Lock()
	defer s.lock.Unlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	if core_utils.IsEmptyOrNil(handle) {
		return "", errors.New("handle is empty")
	}
	claims := make(models.Claims)
	_, err := json.Marshal(info)
	if err != nil {
		return "", err
	}

	now := s.Now()
	createAt := now
	expiresAt := createAt.Add(time.Hour * 24 * 30)

	standardClaims := &jwt.StandardClaims{
		IssuedAt:  createAt.Unix(),
		ExpiresAt: expiresAt.Unix(),
		Issuer:    info.Metadata.Issuer,
		Subject:   info.Metadata.Subject,
		Id:        xid.New().String(),
	}
	var standard map[string]interface{}
	standardJSON, _ := json.Marshal(standardClaims)
	json.Unmarshal(standardJSON, &standard)

	claims["token_info"] = info

	finalClaims := buildClaimsMap(ctx, standardClaims, &claims)

	token, err := s.JWTTokenStore.MintToken(ctx, finalClaims)
	if err != nil {
		return "", err
	}

	return token, nil
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
		return nil, nil
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
