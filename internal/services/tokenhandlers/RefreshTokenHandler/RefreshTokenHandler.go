package RefreshTokenHandler

import (
	"context"
	contracts_clients "echo-starter/internal/contracts/stores/clients"
	contracts_stores_tokenstore "echo-starter/internal/contracts/stores/tokenstore"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	"echo-starter/internal/models"
	"echo-starter/internal/utils"
	"echo-starter/internal/wellknown"
	"encoding/json"
	"net/http"
	"time"

	di "github.com/dozm/di"
	"github.com/fatih/structs"
	core_utils "github.com/fluffy-bunny/fluffycore/utils"
	"github.com/go-oauth2/oauth2/v4/errors"
)

type (
	service struct {
		TokenExchangeTokenHandler     contracts_tokenhandlers.ITokenExchangeTokenHandler     `inject:""`
		ClientCredentialsTokenHandler contracts_tokenhandlers.IClientCredentialsTokenHandler `inject:""`
		ReferenceTokenStore           contracts_stores_tokenstore.ITokenStore                `inject:""`
		ClientRequest                 contracts_clients.IClientRequest                       `inject:""`
	}
	validated struct {
		scopes []string
	}
)

var stemService *service = new(service)

func init() {
	var _ contracts_tokenhandlers.IClientCredentialsTokenHandler = (*service)(nil)
}

func (s *service) Ctor(
	tokenExchangeTokenHandler contracts_tokenhandlers.ITokenExchangeTokenHandler,
	clientCredentialsTokenHandler contracts_tokenhandlers.IClientCredentialsTokenHandler,
	referenceTokenStore contracts_stores_tokenstore.ITokenStore,
	clientRequest contracts_clients.IClientRequest) (*service, error) {
	return &service{
		TokenExchangeTokenHandler:     tokenExchangeTokenHandler,
		ClientCredentialsTokenHandler: clientCredentialsTokenHandler,
		ReferenceTokenStore:           referenceTokenStore,
		ClientRequest:                 clientRequest,
	}, nil
}

// AddScopedIRefreshTokenHandler registers the *service.
func AddScopedIRefreshTokenHandler(builder di.ContainerBuilder) {
	di.AddScoped[contracts_tokenhandlers.IRefreshTokenHandler](builder, stemService.Ctor)
}

func (s *service) ValidationTokenRequest(r *http.Request) (result *contracts_tokenhandlers.ValidatedTokenRequestResult, err error) {
	validated := &contracts_tokenhandlers.ValidatedTokenRequestResult{
		GrantType: r.FormValue("grant_type"),
		Params:    make(map[string]string),
	}
	var safeAddParam = func(key string) {
		val := utils.TrimLeftAndRight(r.FormValue(key))
		if !core_utils.IsEmptyOrNil(val) {
			validated.Params[key] = val
		}
	}
	safeAddParam("scope")
	safeAddParam(models.TokenTypeRefreshToken)

	return validated, nil
}
func (s *service) ProcessTokenRequest(ctx context.Context, result *contracts_tokenhandlers.ValidatedTokenRequestResult) (models.IClaims, error) {
	now := time.Now()
	handle := result.Params[models.TokenTypeRefreshToken]
	rt, err := s.ReferenceTokenStore.GetToken(ctx, handle)
	if err != nil {
		return nil, errors.ErrInvalidRequest
	}
	if rt == nil {
		return nil, errors.ErrInvalidRequest
	}
	if rt.Metadata.Type != models.TokenTypeRefreshToken {
		return nil, errors.ErrInvalidRequest
	}
	if rt.Metadata.ClientID != result.ClientID {
		return nil, errors.New("client_id mismatch")
	}
	// if no scope is passed then we use the scope from the last run
	scope, ok := result.Params["scope"]
	rtInfo := &models.RefreshTokenInfo{}
	rtData, _ := json.Marshal(rt.Data)
	err = json.Unmarshal(rtData, &rtInfo)
	if err != nil {
		return nil, err
	}
	/*
		err = mapstructure.Decode(rt.Data, rtInfo)
		if err != nil {
			return nil, err
		}
	*/
	result.Params = rtInfo.Params
	if ok {
		// override the sone passed into the refresh_token request
		result.Params["scope"] = scope
	}
	newValidatedResult := &contracts_tokenhandlers.ValidatedTokenRequestResult{
		GrantType: rtInfo.GrantType,
		ClientID:  rtInfo.ClientID,
		Params:    result.Params,
	}
	client := s.ClientRequest.GetClient()
	refreshHandle := handle
	if client.RefreshTokenUsage == models.OneTimeOnly {
		handle = utils.GenerateHandle()
	}
	newValidatedResult.RefreshTokenHandle = handle

	var (
		claims models.IClaims
	)
	switch rtInfo.GrantType {
	case wellknown.OAuth2GrantType_ClientCredentials:
		claims, err = s.ClientCredentialsTokenHandler.ProcessTokenRequest(ctx, newValidatedResult)
	case wellknown.OAuth2GrantType_TokenExchange:
		claims, err = s.TokenExchangeTokenHandler.ProcessTokenRequest(ctx, newValidatedResult)

	default:
		return nil, errors.ErrUnsupportedGrantType
	}
	result.GrantType = rtInfo.GrantType
	if err != nil {
		return nil, err
	}
	for {
		if client.RefreshTokenUsage == models.OneTimeOnly && client.RefreshTokenGraceEnabled == false {
			// revoke the old token
			s.ReferenceTokenStore.RemoveToken(ctx, refreshHandle)
			break
		}
		if client.RefreshTokenGraceEnabled == true {
			rtInfo.RefreshTokenGraceAttempts += 1
			if rtInfo.RefreshTokenGraceAttempts >= rtInfo.RefreshTokenGraceMaxAttempts {
				s.ReferenceTokenStore.RemoveToken(ctx, refreshHandle)
				break
			}
			expiration := rt.Metadata.IssedAt.Add(rtInfo.RefreshTokenGraceTTL)
			if now.After(expiration) {
				s.ReferenceTokenStore.RemoveToken(ctx, refreshHandle)
				break
			}

			data := structs.Map(rtInfo)

			rt.Data = data
			s.ReferenceTokenStore.UpdateToken(ctx, refreshHandle, rt)
			break
		}
		break
	}

	result.RefreshTokenHandle = handle
	return claims, nil
}
