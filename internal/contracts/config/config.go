package config

import (
	"reflect"

	"github.com/go-oauth2/oauth2/v4"
)

const (
	Environment_Development = "Development"
)

type (
	RedisOptions struct {
		// The network type, either tcp or unix.
		// Default is tcp.
		Network string `json:"network" mapstructure:"NETWORK"`
		// host:port address.
		Addr string `json:"addr" mapstructure:"ADDR"`
		// Use the specified Username to authenticate the current connection
		// with one of the connections defined in the ACL list when connecting
		// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
		Username string `json:"username" mapstructure:"USERNAME"`
		// Optional password. Must match the password specified in the
		// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
		// or the User Password when connecting to a Redis 6.0 instance, or greater,
		// that is using the Redis ACL system.
		Password string `json:"password" mapstructure:"PASSWORD"`

		Namespace []string `json:"namespace" mapstructure:"NAMESPACE"`
	}
	oidcConfig struct {
		Domain       string `json:"domain" mapstructure:"DOMAIN"`
		ClientID     string `json:"client_id" mapstructure:"CLIENT_ID"`
		ClientSecret string `json:"client_secret" mapstructure:"CLIENT_SECRET" redact:"true"`
		CallbackURL  string `json:"callback_url" mapstructure:"CALLBACK_URL"`
	}
	oauth2Config struct {
		// ClientID is the application's ID.
		ClientID string `json:"client_id" mapstructure:"CLIENT_ID"`

		// ClientSecret is the application's secret.
		ClientSecret string `json:"client_secret" mapstructure:"CLIENT_SECRET" redact:"true"`

		// RedirectURL is the URL to redirect users going through
		// the OAuth flow, after the resource owner's URLs.
		RedirectURL string `json:"redirect_url" mapstructure:"REDIRECT_URL"`

		// Scope specifies optional requested permissions.
		Scopes []string `json:"scopes" mapstructure:"SCOPES"`
	}
	// Config type
	Config struct {
		ApplicationName         string       `json:"applicationName" mapstructure:"APPLICATION_NAME"`
		ApplicationEnvironment  string       `json:"applicationEnvironment" mapstructure:"APPLICATION_ENVIRONMENT"`
		PrettyLog               bool         `json:"prettyLog" mapstructure:"PRETTY_LOG"`
		LogLevel                string       `json:"logLevel" mapstructure:"LOG_LEVEL"`
		Port                    int          `json:"port" mapstructure:"PORT"`
		OIDC                    oidcConfig   `json:"oidc" mapstructure:"OIDC"`
		OAuth2                  oauth2Config `json:"oauth2" mapstructure:"OAUTH2"`
		SessionMaxAgeSeconds    int          `json:"sessionMaxAgeSeconds" mapstructure:"SESSION_MAX_AGE_SECONDS"`
		AuthCookieExpireSeconds int          `json:"authCookieExpireSeconds" mapstructure:"AUTH_COOKIE_EXPIRE_SECONDS"`
		AuthCookieName          string       `json:"authCookieName" mapstructure:"AUTH_COOKIE_NAME"`
		// session|cookie
		AuthStore                 string `json:"authStore" mapstructure:"AUTH_STORE"`
		SecureCookieHashKey       string `json:"secureCookieHashKey" mapstructure:"SECURE_COOKIE_HASH_KEY" redact:"true"`
		SecureCookieEncryptionKey string `json:"secureCookieEncryptionKey" mapstructure:"SECURE_COOKIE_ENCRYPTION_KEY" redact:"true"`
		GraphQLEndpoint           string `json:"graphQLEndpoint" mapstructure:"GRAPHQL_ENDPOINT"`
		// cookie|inmemory|redis
		SessionEngine string `json:"sessionEngine" mapstructure:"SESSION_ENGINE"`
		RedisUrl      string `json:"redisUrl" mapstructure:"REDIS_URL"`
		RedisPassword string `json:"redisPassword" mapstructure:"REDIS_PASSWORD"`

		// github,oidc
		AuthProvider string `json:"authProvider" mapstructure:"AUTH_PROVIDER"`
		SigningKeys  string `json:"signingKeys" mapstructure:"SIGNING_KEYS" redact:"true"`

		ClientStoreProvider             string             `json:"clientStoreProvider" mapstructure:"CLIENT_STORE_PROVIDER"`
		TokenStoreProvider              string             `json:"tokenStoreProvider" mapstructure:"TOKEN_STORE_PROVIDER"`
		AllowedGrantTypes               []oauth2.GrantType `json:"allowedGrantTypes" mapstructure:"ALLOWED_GRANT_TYPES"`
		TokenType                       string             `json:"tokenType" mapstructure:"TOKEN_TYPE"`
		RedisUseMiniRedis               bool               `json:"redisUseMiniRedis" mapstructure:"REDIS_USE_MINIREDIS"`
		RedisOptionsReferenceTokenStore RedisOptions       `json:"redisOptionsReferenceTokenStore" mapstructure:"REDIS_OPTIONS_REFERENCE_TOKEN_STORE"`
		RedisOptionsRefreshTokenStore   RedisOptions       `json:"redisOptionsRefreshTokenStore" mapstructure:"REDIS_OPTIONS_REFRESH_TOKEN_STORE"`
	}
)

var (
	ReflectConfigType = reflect.TypeOf((*Config)(nil))
	// ConfigDefaultJSON default json
	ConfigDefaultJSON = []byte(`
{
	"APPLICATION_NAME": "in-environment",
	"APPLICATION_ENVIRONMENT": "in-environment",
	"PRETTY_LOG": false,
	"LOG_LEVEL": "info",
	"PORT": 1111,
	"OIDC": {
		"DOMAIN": "blah.auth0.com",
		"CLIENT_ID": "in-environment",
		"CLIENT_SECRET": "in-environment",
		"CALLBACK_URL": ""
	},
	"OAUTH2": {
		"CLIENT_ID": "in-environment",
		"CLIENT_SECRET": "in-environment",
		"REDIRECT_URL": "",
		"SCOPES": ""
	},
 	"SESSION_MAX_AGE_SECONDS": 60,
    "AUTH_PROVIDER": "oidc",
	"AUTH_COOKIE_EXPIRE_SECONDS": 60,
	"AUTH_COOKIE_NAME": "_auth",
	"AUTH_STORE": "cookie",
	"SECURE_COOKIE_HASH_KEY": "",
	"SECURE_COOKIE_ENCRYPTION_KEY": "",
	"GRAPHQL_ENDPOINT": "https://countries.trevorblades.com/",
	"SESSION_ENGINE": "cookie",
	"REDIS_URL": "localhost:6379",
	"REDIS_PASSWORD": "",
	"CLIENT_STORE_PROVIDER": "inmemory",
	"TOKEN_STORE_PROVIDER": "inmemory",
	"SIGNING_KEYS": "",
	"ALLOWED_GRANT_TYPES": "client_credentials,refresh_token,urn:ietf:params:oauth:grant-type:token-exchange",
	"TOKEN_TYPE": "Bearer",
	"REDIS_USE_MINIREDIS": false,
	"REDIS_OPTIONS_REFERENCE_TOKEN_STORE": {
		"NETWORK": "tcp",
		"ADDR": "localhost:6379",
		"NAMESPACE": "a,b,c",
		"USERNAME": "",
		"PASSWORD": ""
	},
	"REDIS_OPTIONS_REFRESH_TOKEN_STORE": {
		"NETWORK": "tcp",
		"ADDR": "localhost:6379",
		"NAMESPACE": "a,b,c",
		"USERNAME": "",
		"PASSWORD": ""

	}

}
`)
)
