package domain

type Limiter interface {
	Allow(key string) bool
	Close()
}
