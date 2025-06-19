package agent

import (
	"agent-assigner/internal/dto"
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

func (h *handler) WorkerChatAssign(ctx context.Context, t *asynq.Task) error {
	var (
		data dto.PayloadChatAssign
	)

	err := json.Unmarshal(t.Payload(), &data)
	if err != nil {
		return err
	}

	err = h.service.AgentAssignment(ctx, data.RoomID)
	if err != nil {
		maxRetry, _ := asynq.GetMaxRetry(ctx)
		retryCount, _ := asynq.GetRetryCount(ctx)
		log.Printf("worker task assign: %v, max retry: %d, retry count: %d", err, maxRetry, retryCount)

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
