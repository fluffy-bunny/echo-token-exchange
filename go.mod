module echo-starter

go 1.20

//replace github.com/fluffy-bunny/grpcdotnetgo => ../grpcdotnetgo

//replace github.com/fluffy-bunny/go-redis-search => ../go-redis-search

//replace github.com/fluffy-bunny/rejonson/v8 => ../rejonson
replace github.com/hibiken/asynq => github.com/fluffy-bunny/asynq v0.23.2

replace github.com/dozm/di => github.com/fluffy-bunny/dozm-di v0.3.3

//replace github.com/fluffy-bunny/fluffycore => ../fluffycore

require (
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/ahmetb/go-linq/v3 v3.2.0
	github.com/alexedwards/argon2id v0.0.0-20211130144151-3585854a6387
	github.com/alicebob/miniredis/v2 v2.20.0
	github.com/catmullet/go-workers v1.4.1
	github.com/cheekybits/genny v1.0.0
	github.com/containerd/containerd v1.6.1
	github.com/coreos/go-oidc/v3 v3.1.0
	github.com/dozm/di v0.3.0
	github.com/fatih/structs v1.1.0
	github.com/fluffy-bunny/fluffycore v0.0.2
	github.com/fluffy-bunny/go-redis-search v0.0.5
	github.com/fluffy-bunny/grpcdotnetgo v0.1.252
	github.com/fluffy-bunny/rejonson/v8 v8.0.2
	github.com/fluffy-bunny/sarulabsdi v0.1.65
	github.com/go-oauth2/oauth2/v4 v4.5.2
	github.com/go-oauth2/redis/v4 v4.1.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-test/deep v1.1.0
	github.com/gogo/status v1.1.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/mock v1.6.0
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/jinzhu/copier v0.3.5
	github.com/json-iterator/go v1.1.12
	github.com/labstack/echo-contrib v0.15.0
	github.com/labstack/echo/v4 v4.10.2
	github.com/lestrrat-go/jwx/v2 v2.0.9
	github.com/quasoft/memstore v0.0.0-20191010062613-2bce066d2b0b
	github.com/rbcervilla/redisstore/v8 v8.1.0
	github.com/reugn/async v0.5.0
	github.com/rs/xid v1.5.0
	github.com/rs/zerolog v1.29.1
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.16.0
	github.com/stretchr/testify v1.8.4
	github.com/theTardigrade/golang-passwordGenerator v1.1.1
	github.com/tufin/asciitree v0.0.0-20210127111056-bf70173ef677
	github.com/valyala/fasthttp v1.47.0
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/net v0.10.0
	golang.org/x/oauth2 v0.8.0
	google.golang.org/grpc v1.55.0
)

require (
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/Microsoft/hcsshim v0.9.2 // indirect
	github.com/alicebob/gopher-json v0.0.0-20200520072559-a9ecdc9d1d3a // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/containerd/cgroups v1.0.3 // indirect
	github.com/containerd/continuity v0.2.2 // indirect
	github.com/containerd/fifo v1.0.0 // indirect
	github.com/containerd/ttrpc v1.1.0 // indirect
	github.com/containerd/typeurl v1.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/felixge/fgprof v0.9.3 // indirect
	github.com/fluffy-bunny/viperEx v0.0.27 // indirect
	github.com/form3tech-oss/jwt-go v3.2.5+incompatible // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/pprof v0.0.0-20230323073829-e72429f035bd // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jedib0t/go-pretty/v6 v6.4.6 // indirect
	github.com/jellydator/ttlcache/v2 v2.11.1 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/lestrrat-go/blackmagic v1.0.1 // indirect
	github.com/lestrrat-go/httpcc v1.0.1 // indirect
	github.com/lestrrat-go/httprc v1.0.4 // indirect
	github.com/lestrrat-go/iter v1.0.2 // indirect
	github.com/lestrrat-go/option v1.0.1 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/sys/mountinfo v0.5.0 // indirect
	github.com/moby/sys/signal v0.6.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.1.0 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417 // indirect
	github.com/opencontainers/selinux v1.10.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/profile v1.7.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/tidwall/btree v0.0.0-20191029221954-400434d76274 // indirect
	github.com/tidwall/buntdb v1.1.2 // indirect
	github.com/tidwall/gjson v1.12.1 // indirect
	github.com/tidwall/grect v0.0.0-20161006141115-ba9a043346eb // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/rtree v0.0.0-20180113144539-6cd427091e0e // indirect
	github.com/tidwall/tinyqueue v0.0.0-20180302190814-1e39f5511563 // indirect
	github.com/tkuchiki/go-timezone v0.2.2 // indirect
	github.com/tkuchiki/parsetime v0.3.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/yuin/gopher-lua v0.0.0-20210529063254-f4c35e4016d9 // indirect
	github.com/ziflex/lecho v1.2.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
