package factory

import (
	"agent-assigner/database"
	"agent-assigner/pkg/util"
	"context"
	"fmt"
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

// setAsynqRedisOpt is a function to set asynq redis option
func (f *Factory) setAsynqRedisOpt() {
	opt := database.GetRedisOpt()

	client := asynq.NewClient(opt)
	defer client.Close()

	// Attempt a dummy task enqueue as a health check
	_, err := client.Enqueue(asynq.NewTask("debug:ping", nil), asynq.MaxRetry(0))
	if err != nil {
		fmt.Printf("[REDIS] ❌ Failed to connect to Redis at %s (DB: %d): %v\n", opt.Addr, opt.DB, err)
	} else {
		fmt.Printf("[REDIS] ✅ Connected to Redis at %s (DB: %d)\n", opt.Addr, opt.DB)
	}

	f.asynqRedisOpt = opt
}

// setAsynqWorker is a function to set asynq worker, using for asynq consumer
func (f *Factory) setAsynqWorker(queues map[string]int) {
	// num of worker to process tasks concurrently
	worker, _ := strconv.Atoi(util.GetEnv("ASYNQ_CONCURRENCY", "1"))

	// retry delay on each queue when facing failure
	retryDelay, _ := strconv.Atoi(util.GetEnv("ASYNQ_RETRY_DELAY", "5"))

	// setup async server with configured options
	f.AsynqWorker = asynq.NewServer(f.asynqRedisOpt, asynq.Config{
		Concurrency: worker,
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
