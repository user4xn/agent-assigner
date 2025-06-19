package agent

import (
	"agent-assigner/internal/dto"
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

// this worker will retry 3 time per 5s while facing error
func (h *handler) WorkerChatAssign(ctx context.Context, t *asynq.Task) error {
	var (
		data dto.PayloadChatAssign
	)

	// unmarshal payload into struct
	err := json.Unmarshal(t.Payload(), &data)
	if err != nil {
		return err
	}

	// passing to service layer to working in logic
	err = h.service.AgentAssignment(ctx, data.RoomID)
	if err != nil {
		maxRetry, _ := asynq.GetMaxRetry(ctx)
		retryCount, _ := asynq.GetRetryCount(ctx)
		log.Printf("worker task assign: %v, max retry: %d, retry count: %d", err, maxRetry, retryCount)

		// if max retry reached, re-enqueue room id in 45s
		if retryCount == maxRetry {
			log.Println("max retry reached, re-enqueue room id in 45s: ", data.RoomID)

			err := h.service.RePublishSingleQueue(ctx, data)
			if err != nil {
				log.Println("worker task assign (re-queue): ", err)
				return err
			}
		}

		return err
	}

	log.Println("worker task assign success room id: ", data.RoomID)
	return nil
}
