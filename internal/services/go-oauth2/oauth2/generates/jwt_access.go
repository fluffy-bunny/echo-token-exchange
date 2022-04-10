package generates

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	echo_oauth2 "echo-starter/internal/services/go-oauth2/oauth2"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

// Token based on the UUID generated token
func (a *JWTAccessGenerate) Token(ctx context.Context, data *echo_oauth2.GenerateBasic, isGenRefresh bool) (string, string, error) {
	clientID := data.Client.ClientID
	scopes := strings.Split(data.TokenInfo.GetScope(), " ")
	var scopeInterface interface{}
	if len(scopes) > 1 {
		scopeInterface = scopes
	} else if len(scopes) == 1 {
		scopeInterface = scopes[0]
	}

	claims := &JWTAccessClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        xid.New().String(),
			Audience:  clientID,
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
	token := jwt.NewWithClaims(a.SignedMethod, mClaims)
	if a.SignedKeyID != "" {
		token.Header["kid"] = a.SignedKeyID
	}
	var key interface{}
	if a.isEs() {
		v, err := jwt.ParseECPrivateKeyFromPEM(a.SignedKey)
		if err != nil {
			return "", "", err
		}
		key = v
	} else if a.isRsOrPS() {
		v, err := jwt.ParseRSAPrivateKeyFromPEM(a.SignedKey)
		if err != nil {
			return "", "", err
		}
		key = v
	} else if a.isHs() {
		key = a.SignedKey
	} else {
		return "", "", errors.New("unsupported sign method")
	}

	access, err := token.SignedString(key)
	if err != nil {
		return "", "", err
	}
	refresh := ""

	if isGenRefresh {
		t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(access)).String()
		refresh = base64.URLEncoding.EncodeToString([]byte(t))
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}

	return access, refresh, nil
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
