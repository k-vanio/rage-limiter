package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/k-vanio/rage-limiter/internal/domain"
)

type App struct {
	Persist         domain.Persist
	LimiterPerIp    domain.Limiter
	LimiterPerToken domain.Limiter
}

func New(persist domain.Persist, limiterPerIp, limiterPerToken domain.Limiter) *App {
	return &App{
		Persist:         persist,
		LimiterPerIp:    limiterPerIp,
		LimiterPerToken: limiterPerToken,
	}
}

func (a *App) Run() error {
	mux := http.NewServeMux()

	// handler for / endpoint
	mux.HandleFunc("/", a.rangeLimiter(func(w http.ResponseWriter, r *http.Request) {
		var key string = r.RemoteAddr
		if r.Header.Get("API_KEY") != "" {
			key = r.Header.Get("API_KEY")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rows := a.Persist.Info(key)
		json.NewEncoder(w).Encode(rows)
	}))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return srv.ListenAndServe()
}

// middleware for / endpoint with limiter request
func (a *App) rangeLimiter(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var apiKey string = r.Header.Get("API_KEY")
		if apiKey != "" {
			log.Println("TOKEN: ", apiKey)
			if !a.LimiterPerToken.Allow(apiKey) {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				a.log(apiKey, "you have reached the maximum number of requests or actions allowed within a certain time frame")
				return
			}
		} else {
			apiKey = strings.SplitN(r.RemoteAddr, ":", 2)[0]
			log.Println("IP: ", apiKey)
			if !a.LimiterPerIp.Allow(apiKey) {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				a.log(r.RemoteAddr, "you have reached the maximum number of requests or actions allowed within a certain time frame")
				return
			}
		}

		a.log(apiKey, r.Header)

		fn(w, r)
	}
}

func (a *App) log(key string, data interface{}) {
	err := a.Persist.Store(key, time.Now(), data)
	if err != nil {
		log.Println(err)
	}
}
