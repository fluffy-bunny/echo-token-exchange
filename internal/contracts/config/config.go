package config

import (
	"reflect"

	contracts_jwtvalidator "echo-starter/internal/contracts/jwtvalidator"

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

	// Config type
	Config struct {
		ApplicationName         string `json:"applicationName" mapstructure:"APPLICATION_NAME"`
		ApplicationEnvironment  string `json:"applicationEnvironment" mapstructure:"APPLICATION_ENVIRONMENT"`
		PrettyLog               bool   `json:"prettyLog" mapstructure:"PRETTY_LOG"`
		LogLevel                string `json:"logLevel" mapstructure:"LOG_LEVEL"`
		Port                    int    `json:"port" mapstructure:"PORT"`
		SessionMaxAgeSeconds    int    `json:"sessionMaxAgeSeconds" mapstructure:"SESSION_MAX_AGE_SECONDS"`
		AuthCookieExpireSeconds int    `json:"authCookieExpireSeconds" mapstructure:"AUTH_COOKIE_EXPIRE_SECONDS"`
		AuthCookieName          string `json:"authCookieName" mapstructure:"AUTH_COOKIE_NAME"`
		// session|cookie
		AuthStore                 string `json:"authStore" mapstructure:"AUTH_STORE"`
		SecureCookieHashKey       string `json:"secureCookieHashKey" mapstructure:"SECURE_COOKIE_HASH_KEY" redact:"true"`
		SecureCookieEncryptionKey string `json:"secureCookieEncryptionKey" mapstructure:"SECURE_COOKIE_ENCRYPTION_KEY" redact:"true"`
		// cookie|inmemory|redis
		SessionEngine string `json:"sessionEngine" mapstructure:"SESSION_ENGINE"`
		RedisUrl      string `json:"redisUrl" mapstructure:"REDIS_URL"`
		RedisPassword string `json:"redisPassword" mapstructure:"REDIS_PASSWORD"`

		SigningKeys string `json:"signingKeys" mapstructure:"SIGNING_KEYS" redact:"true"`

		ClientStoreProvider string                                     `json:"clientStoreProvider" mapstructure:"CLIENT_STORE_PROVIDER"`
		TokenStoreProvider  string                                     `json:"tokenStoreProvider" mapstructure:"TOKEN_STORE_PROVIDER"`
		AllowedGrantTypes   []oauth2.GrantType                         `json:"allowedGrantTypes" mapstructure:"ALLOWED_GRANT_TYPES"`
		TokenType           string                                     `json:"tokenType" mapstructure:"TOKEN_TYPE"`
		RedisUseMiniRedis   bool                                       `json:"redisUseMiniRedis" mapstructure:"REDIS_USE_MINIREDIS"`
		RedisOptions        RedisOptions                               `json:"redisOptions" mapstructure:"REDIS_OPTIONS"`
		JWTValidatorOptions contracts_jwtvalidator.JWTValidatorOptions `json:"jwtValidatorOptions" mapstructure:"JWT_VALIDATOR_OPTIONS"`
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
 	"SESSION_MAX_AGE_SECONDS": 60,
 	"AUTH_COOKIE_EXPIRE_SECONDS": 60,
	"AUTH_COOKIE_NAME": "_auth",
	"AUTH_STORE": "cookie",
	"SECURE_COOKIE_HASH_KEY": "",
	"SECURE_COOKIE_ENCRYPTION_KEY": "",
	"SESSION_ENGINE": "cookie",
	"REDIS_URL": "localhost:6379",
	"REDIS_PASSWORD": "",
	"CLIENT_STORE_PROVIDER": "inmemory",
	"TOKEN_STORE_PROVIDER": "inmemory",
	"SIGNING_KEYS": "",
	"ALLOWED_GRANT_TYPES": "client_credentials,refresh_token,urn:ietf:params:oauth:grant-type:token-exchange",
	"TOKEN_TYPE": "Bearer",
	"REDIS_USE_MINIREDIS": true,
	"REDIS_OPTIONS": {
		"NETWORK": "tcp",
		"ADDR": "localhost:6379",
		"NAMESPACE": "a,b,c",
		"USERNAME": "",
		"PASSWORD": ""
	},
	"JWT_VALIDATOR_OPTIONS": {
		"CLOCK_SKEW_MINUTES": 5,
		"VALIDATE_SIGNATURE": true,
		"VALIDATE_ISSUER": true
		"ISSUER": "http://localhost:1523/"
	} 
}
`)
)
