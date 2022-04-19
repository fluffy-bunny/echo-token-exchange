package startup

import (
	"context"
	echostarter_auth "echo-starter/internal/auth"
	contracts_config "echo-starter/internal/contracts/config"
	"echo-starter/internal/models"

	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	"github.com/quasoft/memstore"
	"github.com/rs/zerolog/log"

	services_handlers_about "echo-starter/internal/services/handlers/about"
	app_session "echo-starter/internal/session"

	"net/http"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/startup"
	echo_contracts_startup "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/startup"

	"github.com/gorilla/securecookie"

	services_background_taskclient "echo-starter/internal/services/background/taskclient"
	services_background_taskengine "echo-starter/internal/services/background/taskengine"
	services_background_tasks_removetokens "echo-starter/internal/services/background/tasks/removetokens"

	services_auth_cookie_token_store "echo-starter/internal/services/auth/cookie_token_store"
	services_apiresources_inmemory "echo-starter/internal/services/stores/apiresources/inmemory"
	services_clients_clientrequest "echo-starter/internal/services/stores/clients/clientrequest"
	services_clients_inmemory "echo-starter/internal/services/stores/clients/inmemory"

	services_auth_session_token_store "echo-starter/internal/services/auth/session_token_store"

	services_handlers_healthz "echo-starter/internal/services/handlers/healthz"
	services_handlers_ready "echo-starter/internal/services/handlers/ready"
	services_probes_database "echo-starter/internal/services/probes/database"
	services_probes_oidc "echo-starter/internal/services/probes/oidc"

	// OAuth2
	//----------------------------------------------------------------------------------------------------------------------

	services_stores_jwttoken "echo-starter/internal/services/stores/jwttoken"
	services_stores_keymaterial "echo-starter/internal/services/stores/keymaterial"
	services_stores_tokenstore_inmemory "echo-starter/internal/services/stores/tokenstore/inmemory"

	services_tokenhandlers "echo-starter/internal/services/tokenhandlers"
	services_tokenhandlers_ClientCredentialsTokenHandler "echo-starter/internal/services/tokenhandlers/ClientCredentialsTokenHandler"
	services_tokenhandlers_RefreshTokenHandler "echo-starter/internal/services/tokenhandlers/RefreshTokenHandler"
	services_tokenhandlers_TokenExchangeTokenHandler "echo-starter/internal/services/tokenhandlers/TokenExchangeTokenHandler"

	// OIDC/OAUTH2
	//----------------------------------------------------------------------------------------------------------------------
	services_handlers_api_discovery "echo-starter/internal/services/handlers/api/discovery"
	services_handlers_api_discoveryjwks "echo-starter/internal/services/handlers/api/discoveryjwks"
	services_handlers_api_introspect "echo-starter/internal/services/handlers/api/introspect"
	services_handlers_api_revoke "echo-starter/internal/services/handlers/api/revoke"
	services_handlers_api_token "echo-starter/internal/services/handlers/api/token"

	core_contracts_session "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/session"
	core_middleware_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/middleware/claimsprincipal"

	middleware_oauth2client "echo-starter/internal/middleware/oauth2client"

	middleware_claimsprincipal "echo-starter/internal/middleware/claimsprincipal"
	middleware_session "echo-starter/internal/middleware/session"
	middleware_stores "echo-starter/internal/middleware/stores"

	contracts_background_tasks "echo-starter/internal/contracts/background/tasks"
	services_claimsprovider "echo-starter/internal/services/claimsprovider"
	services_handlers_auth_unauthorized "echo-starter/internal/services/handlers/auth/unauthorized"
	services_handlers_error "echo-starter/internal/services/handlers/error"
	services_handlers_home "echo-starter/internal/services/handlers/home"

	"github.com/fluffy-bunny/go-redis-search/ftsearch"
	core_contracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	contracts_cookies "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/cookies"
	core_middleware_session "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/middleware/session"
	di "github.com/fluffy-bunny/sarulabsdi"
	redis "github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	redisstore "github.com/rbcervilla/redisstore/v8"
)

type Startup struct {
	echo_contracts_startup.CommonStartup
	config       *contracts_config.Config
	ctrl         *gomock.Controller
	clients      []models.Client
	apiResources []models.APIResource
	taskEngine   contracts_background_tasks.ITaskEngine
	container    di.Container
}

func assertImplementation() {
	var _ echo_contracts_startup.IStartup = (*Startup)(nil)
}

func NewStartup() echo_contracts_startup.IStartup {
	signingKeys := os.Getenv("SIGNING_KEYS")
	if core_utils.IsEmptyOrNil(signingKeys) {
		data, err := ioutil.ReadFile("./static/secrets/signing-keys.json")
		if err == nil {
			log.Error().Msg("DO NOT USE THIS IN PRODUCTION: Using signing keys from file")
			os.Setenv("SIGNING_KEYS", string(data))
		}
	}

	startup := &Startup{
		config: &contracts_config.Config{},
		ctrl:   gomock.NewController(nil),
	}
	hooks := &echo_contracts_startup.Hooks{
		PostBuildHook:   startup.PostBuildHook,
		PreStartHook:    startup.PreStartHook,
		PreShutdownHook: startup.PreShutdownHook,
	}

	startup.AddHooks(hooks)

	startup.loadTestClients()
	startup.loadApiResources()
	return startup
}
func (s *Startup) PreStartHook(echo *echo.Echo) error {
	err := s._createDevelopmentIndexes()
	if err != nil {
		return err
	}
	s.taskEngine = contracts_background_tasks.GetITaskEngineFromContainer(s.container)
	return s.taskEngine.Start()
}
func (s *Startup) _createDevelopmentIndexes() error {
	if s.config.ApplicationEnvironment != "Development" {
		return nil
	}
	redisOptions := &redis.Options{
		Addr:     s.config.RedisOptionsReferenceTokenStore.Addr,
		Network:  s.config.RedisOptionsReferenceTokenStore.Network,
		Password: s.config.RedisOptionsReferenceTokenStore.Password,
		Username: s.config.RedisOptionsReferenceTokenStore.Username,
	}
	cli := redis.NewClient(redisOptions)
	indexName := "echoTokenStoreIdx"
	var ftSearch *ftsearch.Client
	ftSearch = ftsearch.NewClient(cli)
	create := ftsearch.NewCreate().WithIndex(indexName).OnJSON().
		WithSchema(ftsearch.NewSchema().
			WithIdentifier("$.metadata.type").AsAttribute("type").AttributeType("TEXT")).
		WithSchema(ftsearch.NewSchema().
			WithIdentifier("$.metadata.client_id").AsAttribute("client_id").AttributeType("TEXT")).
		WithSchema(ftsearch.NewSchema().
			WithIdentifier("$.metadata.subject").AsAttribute("subject").AttributeType("TEXT"))

	_, err := ftSearch.ReIndex(context.Background(), indexName, create)
	return err

}
func (s *Startup) PreShutdownHook(echo *echo.Echo) error {
	return s.taskEngine.Stop()
}
func (s *Startup) PostBuildHook(container di.Container) error {
	if s.config.ApplicationEnvironment == "Development" {
		di.Dump(container)
	}
	s.container = container
	return nil
}
func (s *Startup) getSessionStore() sessions.Store {

	hashKey, err := base64.StdEncoding.DecodeString(s.config.SecureCookieHashKey)
	if err != nil {
		panic(err)
	}
	encryptionKey, err := base64.StdEncoding.DecodeString(s.config.SecureCookieEncryptionKey)
	if err != nil {
		panic(err)
	}

	switch s.config.SessionEngine {
	case "cookie":
		store := sessions.NewCookieStore(hashKey, encryptionKey)
		store.Options.Secure = true
		store.Options.HttpOnly = true
		store.Options.SameSite = http.SameSiteStrictMode
		store.Options.MaxAge = s.config.SessionMaxAgeSeconds
		return store
	case "inmemory":
		store := memstore.NewMemStore(hashKey, encryptionKey)
		store.Options.Secure = true
		store.Options.HttpOnly = true
		store.Options.SameSite = http.SameSiteStrictMode
		store.Options.MaxAge = s.config.SessionMaxAgeSeconds
		return store
	case "redis":
		client := redis.NewClient(&redis.Options{
			Addr:     s.config.RedisUrl,
			Password: s.config.RedisPassword,
		})

		// New default RedisStore
		store, err := redisstore.NewRedisStore(context.Background(), client)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create redis store")

		}
		store.Options(sessions.Options{
			Path:     "/",
			MaxAge:   s.config.SessionMaxAgeSeconds,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})
		return store
	default:
		return nil
	}
}
func (s *Startup) RegisterStaticRoutes(e *echo.Echo) error {
	e.Static("/static", "./static")
	return nil
}

func (s *Startup) GetOptions() *startup.Options {
	return &startup.Options{
		Port: s.config.Port,
	}
}

func (s *Startup) GetConfigOptions() *core_contracts.ConfigOptions {
	prettyLog, err := strconv.ParseBool(os.Getenv("PRETTY_LOG"))
	if err != nil {
		prettyLog = false
	}

	return &core_contracts.ConfigOptions{
		RootConfig:             []byte(contracts_config.ConfigDefaultJSON),
		Destination:            s.config,
		LogLevel:               os.Getenv("LOG_LEVEL"),
		PrettyLog:              prettyLog,
		ApplicationEnvironment: os.Getenv("APPLICATION_ENVIRONMENT"),
	}
}
func (s *Startup) addSecureCookieOptions(builder *di.Builder) {
	// map our config to accessor funcs that other services need
	// SECURE COOKIE
	if core_utils.IsEmptyOrNil(s.config.SecureCookieHashKey) {
		fmt.Println("WARNING: SECURE_COOKIE_HASH_KEY must be set for production......")
		key := securecookie.GenerateRandomKey(32)
		encodedString := base64.StdEncoding.EncodeToString(key)
		s.config.SecureCookieHashKey = encodedString
		fmt.Printf("SECURE_COOKIE_HASH_KEY: %v\n", s.config.SecureCookieHashKey)
	}
	if core_utils.IsEmptyOrNil(s.config.SecureCookieEncryptionKey) {
		fmt.Println("WARNING: SECURE_COOKIE_ENCRYPTION_KEY must be set for production......")
		key := securecookie.GenerateRandomKey(32)
		encodedString := base64.StdEncoding.EncodeToString(key)
		s.config.SecureCookieEncryptionKey = encodedString
		fmt.Printf("SECURE_COOKIE_ENCRYPTION_KEY: %v\n", s.config.SecureCookieEncryptionKey)
	}

	contracts_cookies.AddSecureCookieConfigAccessorFunc(builder, func() *contracts_cookies.SecureCookieConfig {
		return &contracts_cookies.SecureCookieConfig{
			SecureCookieHashKey:       s.config.SecureCookieHashKey,
			SecureCookieEncryptionKey: s.config.SecureCookieEncryptionKey,
		}
	})
}

func (s *Startup) addBackgroundTasksHandlers(builder *di.Builder) {
	// Add the engine
	services_background_taskengine.AddSingletonITaskEngine(builder)

	// Add the client use for enqueing tasks
	services_background_taskclient.AddSingletonITaskClient(builder)

	// add all the handlers
	services_background_tasks_removetokens.AddSingletonISingletonTask(builder)
}
func (s *Startup) addAppHandlers(builder *di.Builder) {

	services_handlers_healthz.AddScopedIHandler(builder)
	services_handlers_ready.AddScopedIHandler(builder)
	services_probes_database.AddSingletonIProbe(builder)
	services_probes_oidc.AddSingletonIProbe(builder)

	services_handlers_home.AddScopedIHandler(builder)
	services_handlers_error.AddScopedIHandler(builder)
	services_handlers_about.AddScopedIHandler(builder)

	// OAuth2
	//----------------------------------------------------------------------------------------------------------------------
	switch s.config.TokenStoreProvider {
	case "inmemory":
		services_stores_tokenstore_inmemory.AddSingletonITokenStore(builder)
	case "redis":
	default:
		panic("token store provider not supported")
	}
	switch s.config.ClientStoreProvider {
	case "inmemory":
		services_clients_inmemory.AddSingletonIClientStore(builder, s.clients)
	default:
		panic("client store provider not supported")
	}

	services_tokenhandlers_ClientCredentialsTokenHandler.AddScopedIClientCredentialsTokenHandler(builder)
	services_tokenhandlers_TokenExchangeTokenHandler.AddScopedITokenExchangeTokenHandler(builder)
	services_tokenhandlers_RefreshTokenHandler.AddScopedIRefreshTokenHandler(builder)
	services_tokenhandlers.AddScopedITokenHandlerAccessor(builder)

	services_stores_keymaterial.AddSingletonIKeyMaterial(builder)
	services_stores_jwttoken.AddSingletonIJwtTokenStore(builder)
	services_clients_clientrequest.AddScopedIClientRequest(builder)
	services_apiresources_inmemory.AddSingletonIAPIResources(builder, s.apiResources)

	// OIDC/OAUTH2
	//----------------------------------------------------------------------------------------------------------------------
	services_handlers_api_discovery.AddScopedIHandler(builder)
	services_handlers_api_discoveryjwks.AddScopedIHandler(builder)
	services_handlers_api_token.AddScopedIHandler(builder)
	services_handlers_api_revoke.AddScopedIHandler(builder)
	services_handlers_api_introspect.AddScopedIHandler(builder)

}

func (s *Startup) ConfigureServices(builder *di.Builder) error {
	dst := &contracts_config.Config{}
	core_utils.PrettyPrintRedacted(s.config, dst)

	// add our config as a sigleton object
	di.AddSingletonTypeByObj(builder, s.config)

	// Add our main session accessor func
	core_contracts_session.AddGetSessionFunc(builder, app_session.GetSession)
	core_contracts_session.AddGetSessionStoreFunc(builder, s.getSessionStore)

	// Add our secure cookie configs
	s.addSecureCookieOptions(builder)

	services_handlers_auth_unauthorized.AddScopedIHandler(builder)

	switch s.config.AuthStore {
	case "session":
		services_auth_session_token_store.AddScopedITokenStore(builder)
	default:
		services_auth_cookie_token_store.AddScopedITokenStore(builder) // overrides the session one
	}

	// add our app handlers
	s.addAppHandlers(builder)

	s.addBackgroundTasksHandlers(builder)

	services_claimsprovider.AddSingletonIClaimsProvider(builder)
	return nil
}
func (s *Startup) Configure(e *echo.Echo, root di.Container) error {
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			id := uuid.New()
			return id.String()
		},
	}))
	e.Use(middleware_stores.EnsureClearExpiredStorageItems(s.GetContainer()))
	// DevelopmentMiddlewareUsingClaimsMap adds all the needed claims so that FinalAuthVerificationMiddlewareUsingClaimsMap succeeds
	//e.Use(middleware_claimsprincipal.DevelopmentMiddlewareUsingClaimsMap(echostarter_auth.BuildGrpcEntrypointPermissionsClaimsMap(), true))
	e.Use(middleware_session.EnsureAuthTokenRefresh(s.GetContainer()))
	e.Use(middleware_oauth2client.AuthenticateOAuth2Client(s.GetContainer()))
	e.Use(middleware_claimsprincipal.AuthenticatedSessionToClaimsPrincipalMiddleware(root))
	e.Use(core_middleware_claimsprincipal.FinalAuthVerificationMiddlewareUsingClaimsMap(echostarter_auth.BuildGrpcEntrypointPermissionsClaimsMap(), true))
	// only after we pass auth do we slide out the auth session
	e.Use(core_middleware_session.EnsureSlidingSession(root, app_session.GetAuthSession))

	return nil
}
