# echo starter  


[auth0-golang-web-app](https://github.com/auth0-samples/auth0-golang-web-app/)  
[demo-echo-app](https://github.com/gtongy/demo-echo-app)  
[cookie auth](https://www.sohamkamani.com/golang/session-cookie-authentication/)

## TLDR  

You can run the app live via docker-compose.  I haven't tested it with Auth0 where the login only sends an access_token as a reference (not jwt) and no refresh_token;
So make sure your Auth0 setup delivers a JWT access_token with a refresh_token.

## Docker-Compose

### Secrets
[Create ECDSA Signing Keys](https://github.com/fluffy-bunny/crypto-gen)  
[Development Signing Keys](cmd/server/static/secrets/signing-keys.json)  
These keys are loaded at startup and placed into the ***SIGNING_KEYS*** environment variable.
```go
  signingKeys := os.Getenv("SIGNING_KEYS")
	if core_utils.IsEmptyOrNil(signingKeys) {
		data, err := ioutil.ReadFile("./static/secrets/signing-keys.json")
		if err == nil {
			log.Error().Msg("DO NOT USE THIS IN PRODUCTION: Using signing keys from file")
			os.Setenv("SIGNING_KEYS", string(data))
		}
	}
```
 
```bash
docker-compose pull
docker-compose up
```

Docker-Compose using [Traefik](https://traefik.io/) to do loadbalancing and gives us an url that doesn't have a port.  
naviage to the following [echo-starter](http://echostarter.docker.localhost/)  

## Motivation

[ECHO](https://echo.labstack.com/) is a fantastic base framework to build upon.  This project adds a lot of design patterns found in [asp.net core](https://docs.microsoft.com/en-us/aspnet/core/introduction-to-aspnet-core).  

1. Introduce [depedency injection](https://github.com/fluffy-bunny/sarulabsdi) with SINGLETON, SCOPED and TRANSIENT features  
As with asp.net core, when a request comes in, we have a context.  A scoped container is created and the handler of that request is a registered as SCOPED.  

The [home handler](internal/services/handlers/home/home.go) is an example.  

Notice the following injected SCOPED ClaimsPrincipal object;  

```go
service struct {
  ClaimsPrincipal contracts_core_claimsprincipal.IClaimsPrincipal `inject:"claimsPrincipal"`
}
```

Just like asp.net core, the claims principal is created on each request.  This is usually populated from an auth cookie or a jwt token.  

2. Introduce asp.net style authentication/authorization using middleware.

Here we have a 2 phase pipeline.  First the claims principal is created but no gating action is done.  A downstream middleware only works on claims principals and gates access to paths.  This allows us to introduce any middleware that can create a claims principal from whatever auth scheme.  Cookie auth and JWT auth are 2 well known schemes that can produce a claims principal.  

1. Templates are ECHO recommendations
2. Sessions are ECHO recommendations
3. Cookies are ECHO recommendations

4. Bring in other nice asp.net standard injectable objects and funcs;  
IDistributedCache  
IMemoryCache  
func ContainerAccessor vs aps.net's IServiceProvider  
etc.  

## Session

Echo uses [gorilla sessions](https://github.com/gorilla/sessions).  Currently this kit supports 3 variants [cookie,inmemory,redis]  
There is NO defaults, you must select one of the supported backends.  

```env
SESSION_ENGINE=redis|cookie|inmemory  
```

If you select redis you must also configure redis  

```env
REDIS_URL=localhost:6379
REDIS_PASSWORD=blah
```  

## Auth

Tokens are stored for backend eyes only.  Even when they are stored as cookies, those cookies are encrypted.  

When selecting ```session``` it ***MUST*** backed by a backend sevice like redis.  The reason is that the cookies get HUGE and the cookie version chunks the big ole FAT auth cookie into parts so that we don't go over the 4K limit.  

```env
AUTH_STORE=cookie|session  
```

If  ```AUTH_STORE=session``` the SESSION_ENGINE ***MUST*** not be cookie.

```env
AUTH_STORE=session
SESSION_ENGINE=inmemory|redis
```  

And if ```AUTH_STORE=cookie``` the SESSION_ENGINE can be any supported one.

```env
AUTH_STORE=cookie
SESSION_ENGINE=cookie|inmemory|redis
```  
