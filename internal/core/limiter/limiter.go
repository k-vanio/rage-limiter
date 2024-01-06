package limiter

type agent struct{}

func New() *agent {
	return &agent{}
}

func (l *agent) Allow(key string) bool {
	return true
}
