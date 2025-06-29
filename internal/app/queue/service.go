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

// function to call factory to these repository
func NewService(f *factory.Factory) Service {
	return &service{
		queuer: f.AsynqClient,
	}
}

// function to enqueue task sparately
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
