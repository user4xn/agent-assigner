package agent

import (
	"agent-assigner/internal/app/queue"
	"agent-assigner/internal/client"
	"agent-assigner/internal/dto"
	"agent-assigner/internal/factory"
	"agent-assigner/pkg/util"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hibiken/asynq"
)

type service struct {
	qiscusClient      client.QiscusClientInterface
	queueService      queue.Service
	adminToken        string
	patternChatAssign string
}

type Service interface {
	WebhookAssigment(ctx context.Context) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		qiscusClient:      client.NewQiscusClient(),
		queueService:      queue.NewService(f),
		adminToken:        util.GetEnv("QISCUS_ADMIN_TOKEN", "fallback"),
		patternChatAssign: util.GetEnv("ASYNQ_PATTERN_CHAT_ASSIGNMENT", "chat:assignment"),
	}
}

func (s *service) WebhookAssigment(ctx context.Context) error {
	serveStatus := "unserved"
	limit := int64(50)

	body := dto.BodyAPIChatRoom{
		Limit:       &limit,
		ServeStatus: &serveStatus,
	}

	fetchRoom, err := s.qiscusClient.FetchUnservedRoom(body)
	if err != nil {
		return err
	}

	for _, r := range fetchRoom.CustomerRooms {
		if r.RoomID == nil {
			log.Println(fmt.Printf("%v %s", r, "room id is nil"))
		}

		roomId, err := strconv.Atoi(*r.RoomID)
		if err != nil {
			log.Println(err)
			continue
		}

		enqueueData := dto.PayloadChatAssign{
			RoomID: int64(roomId),
		}

		for i := 0; i < 3; i++ {
			cctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			payload, _ := json.Marshal(enqueueData)
			opts := []asynq.Option{
				asynq.MaxRetry(3),
				asynq.Unique(24 * time.Hour),
				asynq.Queue(s.patternChatAssign),
			}

			err = s.queueService.Enqueue(cctx, s.patternChatAssign, payload, opts...)
			if err == nil {
				break
			}
		}
	}

	return nil
}
