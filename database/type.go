package database

import (
	"fmt"

	"github.com/hibiken/asynq"
)

type (
	redisConfig struct {
		Host string
		Port string
		Pass string
		DB   int
	}
)

func (conf redisConfig) Configure() {
	// configure redis
	redisOpt = asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Password: conf.Pass,
		DB:       conf.DB,
	}
}
