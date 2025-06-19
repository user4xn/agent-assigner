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
	maxCustomerEnv    string
}

type Service interface {
	WebhookAssigment(ctx context.Context) error
	AgentAssignment(ctx context.Context, roomId int64) error
	RePublishSingleQueue(ctx context.Context, enqueueData dto.PayloadChatAssign) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		qiscusClient:      client.NewQiscusClient(),
		queueService:      queue.NewService(f),
		adminToken:        util.GetEnv("QISCUS_ADMIN_TOKEN", "fallback"),
		patternChatAssign: util.GetEnv("ASYNQ_PATTERN_CHAT_ASSIGNMENT", "chat:assignment"),
		maxCustomerEnv:    util.GetEnv("MAX_CUSTOER_PER_AGENT", "2"),
	}
}

// fetch unserved room and publish queue on each room id
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

		// payload for the queue
		enqueueData := dto.PayloadChatAssign{
			RoomID: int64(roomId),
		}

		// if enqueue task failed, retry 3 times
		for i := 0; i < 3; i++ {
			cctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			// convert payload to json
			payload, _ := json.Marshal(enqueueData)
			// set max retry, unique key for avoiding duplicate task, also the pattern
			opts := []asynq.Option{
				asynq.MaxRetry(3),
				asynq.Unique(1 * time.Minute),
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

func (s *service) AgentAssignment(ctx context.Context, roomId int64) error {
	// get setting of max customer per agent from env
	maxCustomerInt, err := strconv.Atoi(s.maxCustomerEnv)
	if err != nil {
		log.Println("error convert max customer env: ", err)
		return err
	}

	// fetch other agent API
	otherAgent, err := s.qiscusClient.FetchOtherAgent(roomId)
	if err != nil {
		log.Println("error fetch other agent: ", err)
		return err
	}

	for _, a := range otherAgent.Agents {
		// if agent is available and not force offline , and current customer count less than max customer, assign to room
		if a.IsAvailable && !a.ForceOffline && a.CurrentCustomerCount < int64(maxCustomerInt) {
			log.Println("assigning room id: ", roomId, " to agent : ", a.Email)

			body := dto.BodyAssignAgent{
				RoomID:  roomId,
				AgentID: a.ID,
			}

			err = s.qiscusClient.AssignAgent(body)
			if err != nil {
				log.Println("error assign agent: ", err)
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("no agent available")
}

func (s *service) RePublishSingleQueue(ctx context.Context, enqueueData dto.PayloadChatAssign) error {
	time.Sleep(45 * time.Second) // delay 45 to prevent duplicate unique task

	for i := 0; i < 3; i++ {
		cctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		payload, _ := json.Marshal(enqueueData)
		opts := []asynq.Option{
			asynq.MaxRetry(3),
			asynq.Unique(1 * time.Minute),
			asynq.Queue(s.patternChatAssign),
		}

		err := s.queueService.Enqueue(cctx, s.patternChatAssign, payload, opts...)
		if err == nil {
			log.Println("success re-enqueue room id: ", enqueueData.RoomID)
			break
		}
	}

	return fmt.Errorf("error re-enqueue room id: %d", enqueueData.RoomID)
}
