module common_service

go 1.14

require (
	github.com/Blank-Xu/sqlx-adapter v0.0.0-20201228073702-4f0fd6b36b22
	github.com/Masterminds/squirrel v1.5.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/casbin/casbin/v2 v2.31.4
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-co-op/gocron v1.6.2
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-redis/cache/v8 v8.4.1
	github.com/go-redis/redis/v8 v8.4.4
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jmoiron/sqlx v1.3.1
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/rs/zerolog v1.26.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/ugorji/go v1.2.4 // indirect
	github.com/ulule/limiter/v3 v3.8.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/text v0.3.6
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/go-redis/cache/v8 v8.4.1 => github.com/mozhemeng/cache/v8 v8.4.3

replace github.com/spf13/afero => github.com/spf13/afero v1.5.1
