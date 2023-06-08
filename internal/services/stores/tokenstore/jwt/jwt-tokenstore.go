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

	contracts_jwtvalidator "echo-starter/internal/contracts/jwtvalidator"

	di "github.com/dozm/di"
	fluffycore_contracts_common "github.com/fluffy-bunny/fluffycore/contracts/common"
	core_hashset "github.com/fluffy-bunny/fluffycore/gods/sets/hashset"
	core_utils "github.com/fluffy-bunny/fluffycore/utils"
	"github.com/gogo/status"
	"github.com/golang-jwt/jwt"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
)

type (
	service struct {
		JWTTokenStore contracts_stores_tokenstore.IJwtTokenStore `inject:""`
		Now           fluffycore_contracts_common.TimeNow        `inject:""`
		JwtValidator  contracts_jwtvalidator.IJwtValidator       `inject:""`
		lock          *sync.RWMutex
	}
	validated struct {
		scopes []string
	}
)

var stemService *service = new(service)

func (s *service) Ctor(
	jwtValidator contracts_jwtvalidator.IJwtValidator,
	now fluffycore_contracts_common.TimeNow,
	jwtTokenStore contracts_stores_tokenstore.IJwtTokenStore) (*service, error) {
	obj := &service{
		Now:           now,
		JwtValidator:  jwtValidator,
		JWTTokenStore: jwtTokenStore,
		lock:          &sync.RWMutex{},
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
		return nil, status.Error(codes.InvalidArgument, "handle is empty")
	}

	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	s.lock.RLock()
	defer s.lock.RUnlock()
	//--~--~--~--~--~-- BARBED WIRE --~--~--~--~--~--~--
	tt, err := s.JwtValidator.ParseTokenRaw(ctx, handle)
	if err != nil {
		return nil, err
	}
	claims, err := tt.AsMap(ctx)
	if err != nil {
		return nil, err
	}
	info, ok := claims["token_info"]
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "token_info not found")
	}
	infoMap, ok := info.(map[string]interface{})
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "token_info is not a map")
	}
	jb, err := json.Marshal(infoMap)
	if err != nil {
		return nil, err
	}
	tokenInfo := &models.TokenInfo{}
	err = json.Unmarshal(jb, tokenInfo)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil

}
func (s *service) UpdateToken(ctx context.Context, handle string, info *models.TokenInfo) error {

	return nil
}
func (s *service) RemoveToken(ctx context.Context, handle string) error {

	return nil
}
func (s *service) RemoveTokenByClientID(ctx context.Context, clientID string) error {

	return nil
}
func (s *service) RemoveTokenBySubject(ctx context.Context, subject string) error {

	return nil
}
func (s *service) RemoveTokenByClientIdAndSubject(ctx context.Context, clientID string, subject string) error {

	return nil
}
func (s *service) RemoveExpired(ctx context.Context) error {

	return nil
}
