package limiter

import (
	"sync"
	"time"
)

type agent struct {
	mu               *sync.Mutex
	Name             string
	maxRequest       int64
	requests         map[string]int64
	requestsLock     map[string]bool
	timeRequest      time.Duration
	timeLock         time.Duration
	cleanRequest     chan string
	cleanRequestLock chan string
	finish           chan struct{}
}

func New(name string, maxRequest int64, timeRequest, timeLock time.Duration) *agent {
	a := &agent{
		mu:               &sync.Mutex{},
		Name:             name,
		maxRequest:       maxRequest,
		requests:         make(map[string]int64),
		requestsLock:     make(map[string]bool),
		timeRequest:      timeRequest,
		timeLock:         timeLock,
		cleanRequest:     make(chan string),
		cleanRequestLock: make(chan string),
		finish:           make(chan struct{}),
	}

	go func() {
	end:
		for {
			select {
			case key := <-a.cleanRequest:
				a.mu.Lock()
				delete(a.requests, key)
				a.mu.Unlock()
			case key := <-a.cleanRequestLock:
				a.mu.Lock()
				delete(a.requestsLock, key)
				a.mu.Unlock()
			case <-a.finish:
				break end
			}

		}
	}()

	return a
}

func (l *agent) Allow(key string) bool {
	if l.isLock(key) {
		return false
	}

	return l.guard(key)
}

func (l *agent) Close() {
	l.finish <- struct{}{}

	defer func() {
		awaitFor := l.timeRequest * l.timeLock
		time.Sleep(awaitFor)

		close(l.cleanRequest)
		close(l.cleanRequestLock)
		close(l.finish)
	}()
}

func (l *agent) isLock(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if value, ok := l.requestsLock[key]; ok {
		return value
	}

	return false
}

func (l *agent) guard(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if value, ok := l.requests[key]; ok {
		if value >= l.maxRequest {
			l.requestsLock[key] = true
			go func() {
				time.Sleep(l.timeLock)
				l.cleanRequestLock <- key
			}()
			return false
		}

		l.requests[key]++
		return true
	}

	l.requests[key] = 1
	go func() {
		time.Sleep(l.timeRequest)
		l.cleanRequest <- key
	}()

	return true
}
