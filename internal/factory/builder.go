package factory

import "agent-assigner/pkg/util"

func (f *Factory) BuildRestFactory() *Factory {
	// set redis
	f.setAsynqRedisOpt()
	f.setAsynqClient()

	return f
}

func (f *Factory) BuildConsumerFactory() *Factory {
	// set asynq
	f.setAsynqRedisOpt()
	f.setAsynqServer()
	f.setAsynqWorker(map[string]int{
		util.GetEnv("ASYNQ_PATTERN_CHAT_ASSIGNMENT", "chat:assignment"): 10,
	})
	f.setAsynqClient()

	return f
}
