package factory

func (f *Factory) BuildRestFactory() *Factory {
	// set redis
	f.setAsynqRedisOpt()
	f.setAsynqClient()

	return f
}
