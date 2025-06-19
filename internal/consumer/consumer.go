package consumer

import (
	"agent-assigner/internal/app/agent"
	"agent-assigner/internal/factory"
	"agent-assigner/pkg/util"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
)

type consumer struct {
	f                 *factory.Factory
	patternChatAssign string
}

type ConsumerInterface interface {
	Init()
	Stop() error
}

func NewConsumer(f *factory.Factory) *consumer {
	return &consumer{
		f:                 f,
		patternChatAssign: util.GetEnv("ASYNQ_PATTERN_CHAT_ASSIGNMENT", "chat:assignment"),
	}
}

func (c *consumer) Init() {
	// start asynq worker
	go func() {
		err := c.f.AsynqWorker.Start(c.f.AsynqServer)
		if err != nil && err != asynq.ErrServerClosed {
			log.Fatalf("failed to start worker " + err.Error())
		}
	}()

	log.Println("Consumer with asynq started...")

	c.f.AsynqServer.HandleFunc(c.patternChatAssign, agent.NewHandler(c.f).WorkerChatAssign)

	log.Println("Consumer chat assignment registered...")

	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// waiting
	<-sigChan
	c.Stop()

	log.Println("Consumer stopped...")

	os.Exit(0)
}

func (c *consumer) Stop() error {
	if c.f.AsynqWorker != nil {
		c.f.AsynqWorker.Stop()
		c.f.AsynqWorker.Shutdown()
	}

	if c.f.AsynqClient != nil {
		c.f.AsynqClient.Close()
	}

	log.Println("Consumer shutdown...")
	return nil
}
