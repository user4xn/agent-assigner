package factory

import (
	"agent-assigner/database"
	"agent-assigner/pkg/util"
	"context"
	"strconv"
	"time"

	"github.com/hibiken/asynq"
)

type Factory struct {
	AsynqServer   *asynq.ServeMux
	AsynqClient   *asynq.Client
	AsynqWorker   *asynq.Server
	asynqRedisOpt asynq.RedisClientOpt
}

func NewFactory(ctx context.Context) *Factory {
	return &Factory{}
}

func (f *Factory) setAsynqRedisOpt() {
	f.asynqRedisOpt = database.GetRedisOpt()
}

// setAsynqWorker is a function to set asynq worker, using for asynq consumer
func (f *Factory) setAsynqWorker(queues map[string]int) {
	worker, _ := strconv.Atoi(util.GetEnv("ASYNQ_CONCURRENCY", "1"))
	retryDelay, _ := strconv.Atoi(util.GetEnv("ASYNQ_RETRY_DELAY", "2"))
	f.AsynqWorker = asynq.NewServer(f.asynqRedisOpt, asynq.Config{
		Concurrency: worker, // Number of workers to process tasks concurrently
		Queues:      queues,
		RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
			return time.Duration(retryDelay) * time.Second
		},
	})
}

// setAsynqServer is a function to set asynq server, using for asynq consumer
func (f *Factory) setAsynqServer() {
	f.AsynqServer = asynq.NewServeMux()
}

// setAsynqClient is a function to set asynq client, using for asynq producer
func (f *Factory) setAsynqClient() {
	f.AsynqClient = asynq.NewClient(&f.asynqRedisOpt)
}
