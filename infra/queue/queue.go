package asynq

import (
	"crypto/tls"
	"elearning/config"

	"github.com/hibiken/asynq"
)

type QueueClient struct {
	*asynq.Client
}

func CreateClient(cfg *config.Environment) *QueueClient {
	var tlsConfig *tls.Config
	if cfg.RedisUseSSL {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	queueClient := asynq.NewClient(asynq.RedisClientOpt{
		Addr:      cfg.RedisURI,
		Password:  cfg.RedisPassword,
		TLSConfig: tlsConfig,
	})
	return &QueueClient{queueClient}
}
