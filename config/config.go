package config

import (
	"errors"
	gEnv "github.com/Netflix/go-env"
)

var (
	ErrInvalidEnv = errors.New("invalid env")
)

type Environment struct {
	Constants
	RunMode                 string `env:"RUN_MODE,required=true"`
	Port                    int    `env:"PORT,required=true"`
	CorsAllowOrigins        string `env:"CORS_ALLOW_ORIGINS,required=true"`
	MysqlHost               string `env:"MYSQL_HOST"`
	MysqlPort               string `env:"MYSQL_PORT"`
	MysqlUserName           string `env:"MYSQL_USER"`
	MysqlPassword           string `env:"MYSQL_PASSWORD"`
	MysqlDatabase           string `env:"MYSQL_DATABASE"`
	MysqlSSLMode            string `env:"MYSQL_SSL_MODE,required=true"`
	MysqlMigrateMode        bool   `env:"MYSQL_MIGRATE"`
	AccessTokenSecretKey    string `env:"ACCESS_TOKEN_SECRET_KEY,required=true"`
	AccessTokenExpireMinute int    `env:"ACCESS_TOKEN_EXPIRE_MINUTE,required=true"`
	RefreshTokenSecretKey   string `env:"REFRESH_TOKEN_SECRET_KEY,required=true"`
	RefreshTokenExpireHour  int    `env:"REFRESH_TOKEN_EXPIRE_HOUR,required=true"`
	RedisURI                string
	RedisPassword           string
	RedisUseSSL             bool `env:"REDIS_USE_SSL,required=true"`

	FrontendUri  string `env:"FRONTEND_URI,required=true"`
	FrontendPort int    `env:"FRONTEND_PORT,required=true"`
}

func Load() (*Environment, error) {
	var environment Environment
	_, err := gEnv.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, err
	}
	return &environment, nil
}
