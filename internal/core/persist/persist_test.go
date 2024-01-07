package persist_test

import (
	"os"
	"testing"
	"time"

	"github.com/k-vanio/rage-limiter/internal/core/persist"
	"github.com/k-vanio/rage-limiter/internal/domain"
)

func TestStore(t *testing.T) {
	persist := persist.NewRedis(os.Getenv("REDIS_ADDR"), "", 0)

	tests := []struct {
		name    string
		key     string
		time    time.Time
		data    interface{}
		limiter domain.Persist
	}{
		{
			name:    "store data",
			key:     "ip_store",
			time:    time.Now(),
			data:    "data",
			limiter: persist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.limiter.Store(tt.key, tt.time, tt.data)
			if err != nil {
				t.Errorf("Expected nil, got %v", err)
			}
		})
	}
}

func TestInfo(t *testing.T) {
	persist := persist.NewRedis(os.Getenv("REDIS_ADDR"), "", 0)

	tests := []struct {
		name    string
		key     string
		limiter domain.Persist
	}{
		{
			name:    "info data",
			key:     "ip_store",
			limiter: persist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := tt.limiter.Info(tt.key)
			if len(rows) == 0 {
				t.Errorf("Expected not empty, got %v", rows)
			}
		})
	}
}
