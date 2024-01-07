package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/k-vanio/rage-limiter/internal/core/limiter"
	"github.com/k-vanio/rage-limiter/internal/core/persist"
	"github.com/k-vanio/rage-limiter/internal/domain"
	"github.com/k-vanio/rage-limiter/internal/infra/app"
)

func main() {
	maxRequestIp, err := strconv.ParseInt(os.Getenv("MAX_REQUEST_IP_PER_SECOND"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	maxRequestToken, err := strconv.ParseInt(os.Getenv("MAX_REQUEST_TOKEN_PER_SECOND"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	timeLockInt, err := strconv.ParseInt(os.Getenv("TIME_LOCK_IN_SECOND"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("maxRequestIp %v\n", maxRequestIp)
	log.Printf("maxRequestToken %v\n", maxRequestToken)

	timeLock := time.Duration(timeLockInt) * time.Second
	log.Printf("timeLock in seconds %v\n", timeLock.Seconds())

	var store domain.Persist = persist.NewRedis(os.Getenv("REDIS_ADDR"), "", 0)
	var limiterPerIp domain.Limiter = limiter.New("IP", maxRequestIp, time.Second, timeLock)
	defer limiterPerIp.Close()

	var limiterPerToken domain.Limiter = limiter.New("IP", maxRequestToken, time.Second, timeLock)
	defer limiterPerIp.Close()

	var application domain.App = app.New(store, limiterPerIp, limiterPerToken)

	err = application.Run()
	if err != nil {
		log.Fatal(err)
	}
}
