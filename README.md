## Common Service

### Include
- authorization (jwt-go)
- assess control (casbin)
- sql (squirrel & sqlx)
- no sql & cache (go-redis)
- i18n (universal-translator)
- scheduler task (gocron)
- swagger api docs (swag)
- config (viper)
- log manage(logrus & lumberjack)
- command (cobra)

### usage
#### build
```
go build
```

#### prepare
including init tables & create root user
```
common_server prepare/init/setup
```

#### run server
```
common_server server/service/api/run
```

#### run scheduler
```
common_server scheduler/task/job
```

#### update swagger api doc
```
swag init -g cmd/server.go
```

### TODO
- add test
