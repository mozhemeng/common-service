package global

import (
	"common_service/pkg/setting"
	"time"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	CasbinSetting   *setting.CasbinSettingS
	RedisSetting    *setting.RedisSettingS
)

func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	err = s.ReadSection("App", &AppSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("JWT", &JWTSetting)
	if err != nil {
		return err
	}
	JWTSetting.Expire *= time.Second

	err = s.ReadSection("Casbin", &CasbinSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Redis", &RedisSetting)
	if err != nil {
		return err
	}

	return nil
}
