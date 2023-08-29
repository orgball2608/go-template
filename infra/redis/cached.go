package redis

import (
	"context"
	"crypto/tls"
	"elearning/config"

	"github.com/go-redis/redis/v8"
)

type Database struct {
	*redis.Client
}

func ConnectRedis(cfg *config.Environment) (*Database, error) {
	var tlsConfig *tls.Config
	if cfg.RedisUseSSL {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:      cfg.RedisURI,
		Password:  cfg.RedisPassword,
		TLSConfig: tlsConfig,
	})
	err := rdb.Ping(context.Background()).Err()
	return &Database{rdb}, err
}
