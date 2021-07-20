# Common Service

## Include
- api (gin)
- authorization (jwt-go)
- assess control (casbin)
- sql (squirrel & sqlx)
- no sql & cache (go-redis)
- i18n (universal-translator)
- scheduler task (gocron)
- swagger api docs (swag)
- rate limit (limiter)
- config (viper)
- log (logrus & lumberjack)
- command (cobra)

## needed
- redis 
- mysql

## usage
### build
```
go build
```

### prepare
including: create tables & create root role & create root user
```
common_server prepare/init/setup
```

### run server
```
common_server server/service/api/run
```

### run scheduler
```
common_server scheduler/task/job
```

### update swagger api doc
```
swag init -g cmd/server.go
```

## TODO
- add test
