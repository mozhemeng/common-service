# Common Service

## Include
- api (gin)
- authorization (jwt-go)
- access control (casbin)
- sql (squirrel & sqlx)
- no sql & cache (go-redis)
- i18n (universal-translator)
- scheduler task (gocron)
- swagger api docs (swag)
- rate limit (limiter)
- config (viper)
- log (zerolog & lumberjack)
- command (cobra)
- rpc server (grpc)

## needed
- redis 
- mysql

## usage
### build
```
go build
```
### create database
```
CREATE DATABASE IF NOT EXISTS `cose` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
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
