package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	RootPassword         string
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	DefaultPageSize      int
	MaxPageSize          int
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
	UploadDocMaxSize     int
	UploadDocAllowExts   []string
	InitTablesSqlPath    string
}

type DatabaseSettingS struct {
	DBType          string
	Username        string
	Password        string
	Host            string
	Port            string
	DBName          string
	Charset         string
	ParseTime       bool
	MultiStatements bool
	MaxIdleConns    int
	MaxOpenConns    int
}

type JWTSettingS struct {
	Secret        string
	Issuer        string
	RefreshExpire time.Duration
	AccessExpire  time.Duration
}

type CasbinSettingS struct {
	ModelFilePath string
}

type RedisSettingS struct {
	Addr string
	DB   int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
