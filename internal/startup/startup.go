package startup

import (
	"context"
	contracts_config "echo-starter/internal/contracts/config"
	"echo-starter/internal/models"
	services_handlers_about "echo-starter/internal/services/handlers/about"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/alicebob/miniredis/v2"
	contracts_startup "github.com/fluffy-bunny/fluffycore/echo/contracts/startup"
	echo_contracts_startup "github.com/fluffy-bunny/fluffycore/echo/contracts/startup"
	echo_services_startup "github.com/fluffy-bunny/fluffycore/echo/services/startup"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	"github.com/quasoft/memstore"
	"github.com/rs/zerolog/log"

	fluffycore_contracts_runtime "github.com/fluffy-bunny/fluffycore/contracts/runtime"

	services_apiresources_inmemory "echo-starter/internal/services/stores/apiresources/inmemory"
	services_clients_clientrequest "echo-starter/internal/services/stores/clients/clientrequest"
	services_clients_inmemory "echo-starter/internal/services/stores/clients/inmemory"

	services_handlers_healthz "echo-starter/internal/services/handlers/healthz"
	services_handlers_ready "echo-starter/internal/services/handlers/ready"
	services_probes_database "echo-starter/internal/services/probes/database"
	services_probes_oidc "echo-starter/internal/services/probes/oidc"

	// OAuth2
	//----------------------------------------------------------------------------------------------------------------------

	services_jwtvalidator "echo-starter/internal/services/jwtvalidator"
	services_stores_jwttoken "echo-starter/internal/services/stores/jwttoken"
	services_stores_keymaterial "echo-starter/internal/services/stores/keymaterial"
	services_stores_tokenstore_inmemory "echo-starter/internal/services/stores/tokenstore/inmemory"
	services_stores_tokenstore_jwt "echo-starter/internal/services/stores/tokenstore/jwt"
	services_stores_tokenstore_rejson "echo-starter/internal/services/stores/tokenstore/rejson"

	services_tokenhandlers "echo-starter/internal/services/tokenhandlers"
	services_tokenhandlers_ClientCredentialsTokenHandler "echo-starter/internal/services/tokenhandlers/ClientCredentialsTokenHandler"
	services_tokenhandlers_RefreshTokenHandler "echo-starter/internal/services/tokenhandlers/RefreshTokenHandler"
	services_tokenhandlers_TokenExchangeTokenHandler "echo-starter/internal/services/tokenhandlers/TokenExchangeTokenHandler"

	// OIDC/OAUTH2
	//----------------------------------------------------------------------------------------------------------------------
	services_handlers_api_discovery "echo-starter/internal/services/handlers/api/discovery"
	services_handlers_api_discoveryjwks "echo-starter/internal/services/handlers/api/discoveryjwks"

	services_handlers_api_token "echo-starter/internal/services/handlers/api/token"

	middleware_oauth2client "echo-starter/internal/middleware/oauth2client"

	services_claimsprovider "echo-starter/internal/services/claimsprovider"

	services_handlers_home "echo-starter/internal/services/handlers/home"

	di "github.com/dozm/di"
	"github.com/fluffy-bunny/go-redis-search/ftsearch"
	redis "github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	redisstore "github.com/rbcervilla/redisstore/v8"
)

type Startup struct {
	echo_services_startup.StartupBase
	config            *contracts_config.Config
	ctrl              *gomock.Controller
	clients           []models.Client
	apiResources      []models.APIResource
	container         di.Container
	miniRedisInstance *miniredis.Miniredis
}

func init() {
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
	if s.config.RedisUseMiniRedis {
		s.miniRedisInstance = miniredis.NewMiniRedis()
		s.miniRedisInstance.RequireAuth(s.config.RedisOptions.Password)

		err := s.miniRedisInstance.Start()
		if err != nil {
			panic(err)
		}
		s.config.RedisOptions.Addr = s.miniRedisInstance.Addr()

	}
	return nil
}
func (s *Startup) _createDevelopmentIndexes() error {
	if s.config.ApplicationEnvironment != "Development" {
		return nil
	}
	redisOptions := &redis.Options{
		Addr:     s.config.RedisOptions.Addr,
		Network:  s.config.RedisOptions.Network,
		Password: s.config.RedisOptions.Password,
		Username: s.config.RedisOptions.Username,
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
	if s.config.RedisUseMiniRedis {
		s.miniRedisInstance.Close()
	}
	return nil
}
func (s *Startup) PostBuildHook(container di.Container) error {
	if s.config.ApplicationEnvironment == "Development" {
		//		di.Dump(container)
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

func (s *Startup) GetOptions() *contracts_startup.Options {
	return &contracts_startup.Options{
		Port: s.config.Port,
	}
}

func (s *Startup) GetConfigOptions() *fluffycore_contracts_runtime.ConfigOptions {

	return &fluffycore_contracts_runtime.ConfigOptions{
		RootConfig:  []byte(contracts_config.ConfigDefaultJSON),
		Destination: s.config,
	}
}

func (s *Startup) addAppHandlers(builder di.ContainerBuilder) {

	services_handlers_healthz.AddScopedIHandler(builder)
	services_handlers_ready.AddScopedIHandler(builder)
	services_probes_database.AddSingletonIProbe(builder)
	services_probes_oidc.AddSingletonIProbe(builder)

	services_handlers_home.AddScopedIHandler(builder)
	services_handlers_about.AddScopedIHandler(builder)

	// OAuth2
	//----------------------------------------------------------------------------------------------------------------------
	switch s.config.TokenStoreProvider {
	case "inmemory":
		services_stores_tokenstore_inmemory.AddSingletonITokenStore(builder)
	case "redis":
		services_stores_tokenstore_rejson.AddSingletonITokenStore(builder)
	case "jwt":
		services_stores_tokenstore_jwt.AddSingletonITokenStore(builder)
	default:
		panic("token store provider not supported")
	}
	switch s.config.ClientStoreProvider {
	case "inmemory":
		services_clients_inmemory.AddSingletonIClientStore(builder, s.clients)
	default:
		panic("client store provider not supported")
	}

	services_jwtvalidator.AddSingletonIJwtValidator(builder)

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

}

func (s *Startup) ConfigureServices(builder di.ContainerBuilder) error {
	dst := &contracts_config.Config{}
	core_utils.PrettyPrintRedacted(s.config, dst)

	// add our config as a sigleton object
	di.AddInstance[*contracts_config.Config](builder, s.config)

	// add our app handlers
	s.addAppHandlers(builder)

	services_claimsprovider.AddSingletonIClaimsProvider(builder)
	return nil
}
func (s *Startup) Configure(e *echo.Echo, root di.Container) error {

	// DevelopmentMiddlewareUsingClaimsMap adds all the needed claims so that FinalAuthVerificationMiddlewareUsingClaimsMap succeeds
	//e.Use(middleware_claimsprincipal.DevelopmentMiddlewareUsingClaimsMap(echostarter_auth.BuildGrpcEntrypointPermissionsClaimsMap(), true))
	e.Use(middleware_oauth2client.AuthenticateOAuth2Client(s.GetContainer()))

	return nil
}
