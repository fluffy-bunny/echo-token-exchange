package generates

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	echo_oauth2 "echo-starter/internal/services/go-oauth2/oauth2"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/golang-jwt/jwt"
	"github.com/rs/xid"
)

var notAllowed *core_hashset.StringSet

func init() {
	notAllowed = core_hashset.NewStringSet()
	notAllowed.Add("exp", "scope", "aud", "iss", "sub", "iat", "nbf", "jti", "client_id")
}

type CustomClaims struct {
	ClientID string      `json:"client_id,omitempty"`
	Scope    interface{} `json:"scope,omitempty"`
}

// JWTAccessClaims jwt claims
type JWTAccessClaims struct {
	jwt.StandardClaims
	CustomClaims
}

type mapClaims map[string]interface{}

// Valid claims verification
func (a mapClaims) Valid() error {
	return nil
}

// Valid claims verification
func (a *JWTAccessClaims) Valid() error {
	if time.Unix(a.ExpiresAt, 0).Before(time.Now()) {
		return errors.ErrInvalidAccessToken
	}
	return nil
}

// NewJWTAccessGenerate create to generate the jwt access token instance
func NewJWTAccessGenerate(kid string, key []byte, method jwt.SigningMethod) *JWTAccessGenerate {
	return &JWTAccessGenerate{
		SignedKeyID:  kid,
		SignedKey:    key,
		SignedMethod: method,
	}
}

// JWTAccessGenerate generate the jwt access token
type JWTAccessGenerate struct {
	SignedKeyID  string
	SignedKey    []byte
	SignedMethod jwt.SigningMethod
	Issuer       string
}

func (a *JWTAccessGenerate) SetIssuer(issuer string) {
	a.Issuer = issuer
}
func IsSliceOrArray(v interface{}) bool {
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		return true
	}
	if reflect.TypeOf(v).Kind() == reflect.Array {
		return true
	}
	return false
}

func SliceOrArrayLength(v interface{}) int {
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		return reflect.ValueOf(v).Len()
	}
	if reflect.TypeOf(v).Kind() == reflect.Array {
		return reflect.ValueOf(v).Len()
	}
	return 0
}

// Token based on the UUID generated token
func (a *JWTAccessGenerate) Token(ctx context.Context, data *echo_oauth2.GenerateBasic,
	extraClaims contracts_tokenhandlers.Claims) (string, error) {
	clientID := data.Client.ClientID
	scopes := strings.Split(data.TokenInfo.GetScope(), " ")
	var scopeInterface interface{}
	if len(scopes) > 1 {
		scopeInterface = scopes
	} else if len(scopes) == 1 {
		scopeInterface = scopes[0]
	}

	// special case, aud is allowed
	audienceSet := core_hashset.NewStringSet(clientID)
	if !core_utils.IsEmptyOrNil(extraClaims) {
		extraAudInterface := extraClaims["aud"]
		fmt.Println(IsSliceOrArray(extraAudInterface))
		fmt.Println(SliceOrArrayLength(extraAudInterface))
		switch extraAudInterface.(type) {
		case string:
			audienceSet.Add(extraAudInterface.(string))
		case []string:
			audienceSet.Add(extraAudInterface.([]string)...)
		}
	}
	// sanitize the extra claims.
	var toBeRemoved []string
	for key := range extraClaims {
		if notAllowed.Contains(key) {
			toBeRemoved = append(toBeRemoved, key)
		}
	}
	for _, key := range toBeRemoved {
		delete(extraClaims, key)
	}

	claims := &JWTAccessClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        xid.New().String(),
			Issuer:    a.Issuer,
			Subject:   data.UserID,
			ExpiresAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
		},
		CustomClaims: CustomClaims{
			ClientID: clientID,
			Scope:    scopeInterface,
		},
	}
	mClaims := make(mapClaims)
	firstClaims, _ := json.Marshal(claims)
	json.Unmarshal(firstClaims, &mClaims)

	mClaims["aud"] = audienceSet.Values()
	// add in all the extraClaims
	for key, value := range extraClaims {
		mClaims[key] = value
	}

	var arbitrary map[string][]string = make(map[string][]string)
	if !core_utils.IsEmptyOrNil(data.Client.Claims) {
		for _, v := range data.Client.Claims {
			_, ok := arbitrary[v.Type]
			if !ok {
				arbitrary[v.Type] = make([]string, 0)
			}
			arbitrary[v.Type] = append(arbitrary[v.Type], v.Value)
		}
	}

	for key, v := range arbitrary {
		if notAllowed.Contains(key) {
			continue
		}
		_, ok := mClaims[key]
		if ok {
			ori := mClaims[key].([]interface{})
			for _, value := range v {
				ori = append(ori, value)
			}
			mClaims[key] = ori
		} else {
			if len(v) > 1 {
				mClaims[key] = v // add all the strings
			} else if len(v) == 1 {
				mClaims[key] = v[0] // add the first string
			}
		}
	}
	// make sure a string array with only one item is just a string in the jwt
	for key, value := range mClaims {
		switch value.(type) {
		case []string:
			if len(value.([]string)) == 1 {
				mClaims[key] = value.([]string)[0]
			}
		}
	}

	token := jwt.NewWithClaims(a.SignedMethod, mClaims)
	if a.SignedKeyID != "" {
		token.Header["kid"] = a.SignedKeyID
	}
	var key interface{}
	if a.isEs() {
		v, err := jwt.ParseECPrivateKeyFromPEM(a.SignedKey)
		if err != nil {
			return "", err
		}
		key = v
	} else if a.isRsOrPS() {
		v, err := jwt.ParseRSAPrivateKeyFromPEM(a.SignedKey)
		if err != nil {
			return "", err
		}
		key = v
	} else if a.isHs() {
		key = a.SignedKey
	} else {
		return "", errors.New("unsupported sign method")
	}

	access, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return access, nil
}

func (a *JWTAccessGenerate) isEs() bool {
	return strings.HasPrefix(a.SignedMethod.Alg(), "ES")
}

func (a *JWTAccessGenerate) isRsOrPS() bool {
	isRs := strings.HasPrefix(a.SignedMethod.Alg(), "RS")
	isPs := strings.HasPrefix(a.SignedMethod.Alg(), "PS")
	return isRs || isPs
}

func (a *JWTAccessGenerate) isHs() bool {
	return strings.HasPrefix(a.SignedMethod.Alg(), "HS")
}
