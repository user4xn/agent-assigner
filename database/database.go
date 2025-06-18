package database

import (
	"agent-assigner/pkg/util"
	"strconv"
	"sync"

	"github.com/hibiken/asynq"
)

var (
	redisOpt asynq.RedisClientOpt
	once     sync.Once
)

func CreateConnection() {
	REDIS_DB, err := strconv.Atoi(util.GetEnv("REDIS_DB", "0"))
	if err != nil {
		panic(err)
	}

	redis := redisConfig{
		Host: util.GetEnv("REDIS_HOST", "localhost"),
		Port: util.GetEnv("REDIS_PORT", "6379"),
		Pass: util.GetEnv("REDIS_PASS", "fallback"),
		DB:   REDIS_DB,
	}

	once.Do(func() {
		redis.Configure()
	})
}

// Check redis connection, if exist return the memory address of the redis connection
func GetRedisOpt() asynq.RedisClientOpt {
	if redisOpt == (asynq.RedisClientOpt{}) {
		CreateConnection()
	}

	return redisOpt
}
