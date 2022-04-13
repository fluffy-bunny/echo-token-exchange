package manage

import (
	"context"
	"time"

	contracts_clients "echo-starter/internal/contracts/clients"
	contracts_tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	echo_models "echo-starter/internal/models"
	echo_oauth2 "echo-starter/internal/services/go-oauth2/oauth2"

	contracts_stores_apiresources "echo-starter/internal/contracts/stores/apiresources"
	contracts_stores_refreshtoken "echo-starter/internal/contracts/stores/refreshtoken"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	oauth2_manage "github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
)

// NewDefaultManager create to default authorization management instance
func NewDefaultManager() *Manager {
	m := NewManager()
	return m
}

// NewManager create to authorization management instance
func NewManager() *Manager {
	return &Manager{
		gtcfg: make(map[oauth2.GrantType]*oauth2_manage.Config),
	}
}

// Manager provide authorization management
type Manager struct {
	gtcfg             map[oauth2.GrantType]*oauth2_manage.Config
	rcfg              *oauth2_manage.RefreshingConfig
	accessGenerate    echo_oauth2.AccessGenerate
	tokenStore        oauth2.TokenStore
	clientStore       contracts_clients.IClientStore
	apiResources      contracts_stores_apiresources.IAPIResources
	RefreshTokenStore contracts_stores_refreshtoken.IRefreshTokenStore
}

// get grant type config
func (m *Manager) grantConfig(gt oauth2.GrantType) *oauth2_manage.Config {
	if c, ok := m.gtcfg[gt]; ok && c != nil {
		return c
	}
	switch gt {
	case oauth2.Refreshing:
		return &oauth2_manage.Config{
			RefreshTokenExp: time.Minute * 5, // TODO config this
		}
	case oauth2.ClientCredentials:
		//		return oauth2_manage.DefaultClientTokenCfg
		return &oauth2_manage.Config{
			AccessTokenExp:    time.Minute * 30,
			IsGenerateRefresh: true, // TODO: client_credentials should NOT have a refresh_token
		}
	}
	return &oauth2_manage.Config{}
}

// SetClientTokenCfg set the client grant token config
func (m *Manager) SetClientTokenCfg(cfg *oauth2_manage.Config) {
	m.gtcfg[oauth2.ClientCredentials] = cfg
}

// SetRefreshTokenCfg set the refreshing token config
func (m *Manager) SetRefreshTokenCfg(cfg *oauth2_manage.RefreshingConfig) {
	m.rcfg = cfg
}

// MapAccessGenerate mapping the access token generate interface
func (m *Manager) MapAccessGenerate(gen echo_oauth2.AccessGenerate) {
	m.accessGenerate = gen
}

// MapClientStorage mapping the client store interface
func (m *Manager) MapClientStorage(stor contracts_clients.IClientStore) {
	m.clientStore = stor
}

// MustClientStorage mandatory mapping the client store interface
func (m *Manager) MustClientStorage(stor contracts_clients.IClientStore, err error) {
	if err != nil {
		panic(err.Error())
	}
	m.clientStore = stor
}

// MustApiResources mandatory mapping the client store interface
func (m *Manager) MustApiResources(stor contracts_stores_apiresources.IAPIResources, err error) {
	if err != nil {
		panic(err.Error())
	}
	m.apiResources = stor
}

// MustRefreshTokenStore mandatory mapping the client store interface
func (m *Manager) MustRefreshTokenStore(stor contracts_stores_refreshtoken.IRefreshTokenStore, err error) {
	if err != nil {
		panic(err.Error())
	}
	m.RefreshTokenStore = stor
}

// MapTokenStorage mapping the token store interface
func (m *Manager) MapTokenStorage(stor oauth2.TokenStore) {
	m.tokenStore = stor
}

// MustTokenStorage mandatory mapping the token store interface
func (m *Manager) MustTokenStorage(stor oauth2.TokenStore, err error) {
	if err != nil {
		panic(err)
	}
	m.tokenStore = stor
}

// GetClient get the client information
func (m *Manager) GetClient(ctx context.Context, clientID string) (cli *echo_models.Client, err error) {
	cli, found, err := m.clientStore.GetClient(ctx, clientID)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.ErrInvalidClient
	}
	return
}

// GenerateAccessToken generate the access token
func (m *Manager) GenerateAccessToken(ctx context.Context,
	validatedResult *contracts_tokenhandlers.ValidatedTokenRequestResult,
	gt oauth2.GrantType, tgr *oauth2.TokenGenerateRequest, claims contracts_tokenhandlers.Claims) (oauth2.TokenInfo, error) {
	cli, err := m.GetClient(ctx, tgr.ClientID)
	if err != nil {
		return nil, err
	}

	ti := models.NewToken()
	ti.SetClientID(tgr.ClientID)
	ti.SetUserID(tgr.UserID)
	ti.SetRedirectURI(tgr.RedirectURI)
	scope, _ := validatedResult.Params["scope"]
	ti.SetScope(scope)

	createAt := time.Now()
	ti.SetAccessCreateAt(createAt)

	// set access token expires
	gcfg := m.grantConfig(gt)
	aexp := gcfg.AccessTokenExp
	if exp := tgr.AccessTokenExp; exp > 0 {
		aexp = exp
	}
	ti.SetAccessExpiresIn(aexp)
	if gcfg.IsGenerateRefresh {
		ti.SetRefreshCreateAt(createAt)
		ti.SetRefreshExpiresIn(gcfg.RefreshTokenExp)
	}

	td := &echo_oauth2.GenerateBasic{
		APIResources: m.apiResources,
		Client:       cli,
		UserID:       tgr.UserID,
		CreateAt:     createAt,
		TokenInfo:    ti,
		Request:      tgr.Request,
	}

	av, err := m.accessGenerate.Token(ctx, td, claims)
	if err != nil {
		return nil, err
	}
	ti.SetAccess(av)

	if gcfg.IsGenerateRefresh {
		var absoluteExpiration = time.Now().Add(time.Second * time.Duration(cli.AbsoluteRefreshTokenLifetime))
		if cli.AbsoluteRefreshTokenLifetime <= 0 {
			absoluteExpiration = time.Now().Add(time.Hour * 24 * 365 * 10) // 10 years
		}
		var expiration = time.Now().Add(time.Second * time.Duration(cli.RefreshTokenExpiration))
		handle, err := m.RefreshTokenStore.StoreRefreshToken(ctx,
			&contracts_stores_refreshtoken.RefreshTokenInfo{
				ClientID:           tgr.ClientID,
				Subject:            tgr.UserID,
				Scope:              tgr.Scope,
				GrantType:          validatedResult.GrantType,
				Expiration:         expiration,
				AbsoluteExpiration: absoluteExpiration,
				Params:             validatedResult.Params,
			})
		if err != nil {
			return nil, err
		}
		ti.SetRefresh(handle)
	}

	err = m.tokenStore.Create(ctx, ti)
	if err != nil {
		return nil, err
	}

	return ti, nil
}

// RefreshAccessToken refreshing an access token
func (m *Manager) RefreshAccessToken(ctx context.Context, tgr *oauth2.TokenGenerateRequest) (oauth2.TokenInfo, error) {
	ti, err := m.LoadRefreshToken(ctx, tgr.Refresh)
	if err != nil {
		return nil, err
	}

	cli, err := m.GetClient(ctx, ti.GetClientID())
	if err != nil {
		return nil, err
	}

	oldAccess, oldRefresh := ti.GetAccess(), ti.GetRefresh()

	td := &echo_oauth2.GenerateBasic{
		Client:    cli,
		UserID:    ti.GetUserID(),
		CreateAt:  time.Now(),
		TokenInfo: ti,
		Request:   tgr.Request,
	}

	rcfg := oauth2_manage.DefaultRefreshTokenCfg
	if v := m.rcfg; v != nil {
		rcfg = v
	}

	ti.SetAccessCreateAt(td.CreateAt)
	if v := rcfg.AccessTokenExp; v > 0 {
		ti.SetAccessExpiresIn(v)
	}

	if v := rcfg.RefreshTokenExp; v > 0 {
		ti.SetRefreshExpiresIn(v)
	}

	if rcfg.IsResetRefreshTime {
		ti.SetRefreshCreateAt(td.CreateAt)
	}

	if scope := tgr.Scope; scope != "" {
		ti.SetScope(scope)
	}

	tv, err := m.accessGenerate.Token(ctx, td, nil)
	if err != nil {
		return nil, err
	}

	ti.SetAccess(tv)
	rv := "" // TODO: FIX
	if rv != "" {
		ti.SetRefresh(rv)
	}

	if err := m.tokenStore.Create(ctx, ti); err != nil {
		return nil, err
	}

	if rcfg.IsRemoveAccess {
		// remove the old access token
		if err := m.tokenStore.RemoveByAccess(ctx, oldAccess); err != nil {
			return nil, err
		}
	}

	if rcfg.IsRemoveRefreshing && rv != "" {
		// remove the old refresh token
		if err := m.tokenStore.RemoveByRefresh(ctx, oldRefresh); err != nil {
			return nil, err
		}
	}

	if rv == "" {
		ti.SetRefresh("")
		ti.SetRefreshCreateAt(time.Now())
		ti.SetRefreshExpiresIn(0)
	}

	return ti, nil
}

// RemoveAccessToken use the access token to delete the token information
func (m *Manager) RemoveAccessToken(ctx context.Context, access string) error {
	if access == "" {
		return errors.ErrInvalidAccessToken
	}
	return m.tokenStore.RemoveByAccess(ctx, access)
}

// RemoveRefreshToken use the refresh token to delete the token information
func (m *Manager) RemoveRefreshToken(ctx context.Context, refresh string) error {
	if refresh == "" {
		return errors.ErrInvalidAccessToken
	}
	return m.tokenStore.RemoveByRefresh(ctx, refresh)
}

// LoadAccessToken according to the access token for corresponding token information
func (m *Manager) LoadAccessToken(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	if access == "" {
		return nil, errors.ErrInvalidAccessToken
	}

	ct := time.Now()
	ti, err := m.tokenStore.GetByAccess(ctx, access)
	if err != nil {
		return nil, err
	} else if ti == nil || ti.GetAccess() != access {
		return nil, errors.ErrInvalidAccessToken
	} else if ti.GetRefresh() != "" && ti.GetRefreshExpiresIn() != 0 &&
		ti.GetRefreshCreateAt().Add(ti.GetRefreshExpiresIn()).Before(ct) {
		return nil, errors.ErrExpiredRefreshToken
	} else if ti.GetAccessExpiresIn() != 0 &&
		ti.GetAccessCreateAt().Add(ti.GetAccessExpiresIn()).Before(ct) {
		return nil, errors.ErrExpiredAccessToken
	}
	return ti, nil
}

// LoadRefreshToken according to the refresh token for corresponding token information
func (m *Manager) LoadRefreshToken(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	if refresh == "" {
		return nil, errors.ErrInvalidRefreshToken
	}

	ti, err := m.tokenStore.GetByRefresh(ctx, refresh)
	if err != nil {
		return nil, err
	} else if ti == nil || ti.GetRefresh() != refresh {
		return nil, errors.ErrInvalidRefreshToken
	} else if ti.GetRefreshExpiresIn() != 0 && // refresh token set to not expire
		ti.GetRefreshCreateAt().Add(ti.GetRefreshExpiresIn()).Before(time.Now()) {
		return nil, errors.ErrExpiredRefreshToken
	}
	return ti, nil
}
