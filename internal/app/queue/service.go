package queue

import (
	"agent-assigner/internal/factory"
	"context"
	"log"

	"github.com/hibiken/asynq"
)

type service struct {
	queuer *asynq.Client
}

type Service interface {
	Enqueue(ctx context.Context, pattern string, payload []byte, opts ...asynq.Option) error
}

// A function to call factory to initialize database connection to this/these repository
func NewService(f *factory.Factory) Service {
	return &service{
		queuer: f.AsynqClient,
	}
}

func (s *service) Enqueue(ctx context.Context, pattern string, payload []byte, opts ...asynq.Option) error {
	task := asynq.NewTask(pattern, payload, opts...)

	_, err := s.queuer.EnqueueContext(ctx, task)
	if err != nil {
		log.Printf("Enqueue task failed: %v", err)

		return err
	}

	log.Println("Enqueue task success")
	return nil
}
