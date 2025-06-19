package factory

import "agent-assigner/pkg/util"

func (f *Factory) BuildRestFactory() *Factory {
	// set redis & asynq producer
	f.setAsynqRedisOpt()
	f.setAsynqClient()

	return f
}

func (f *Factory) BuildConsumerChatAssignFactory() *Factory {
	// set asynq consumer
	f.setAsynqRedisOpt()
	f.setAsynqServer()

	// configure worker and the pattern of consumer
	f.setAsynqWorker(map[string]int{
		util.GetEnv("ASYNQ_PATTERN_CHAT_ASSIGNMENT", "chat:assignment"): 10,
	})

	// still need for re enqueue task
	f.setAsynqClient()

	return f
}
