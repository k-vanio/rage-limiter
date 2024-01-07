package limiter_test

import (
	"testing"
	"time"

	"github.com/k-vanio/rage-limiter/internal/core/limiter"
	"github.com/k-vanio/rage-limiter/internal/domain"
)

func TestAllow(t *testing.T) {
	limiter := limiter.New("ip_any", 10, time.Millisecond, time.Microsecond)
	defer limiter.Close()

	tests := []struct {
		name     string
		key      string
		try      int64
		waitFor  time.Duration
		expected bool
		limiter  domain.Limiter
	}{
		{
			name:     "allow request",
			key:      "ip_allow",
			try:      10,
			waitFor:  time.Microsecond,
			expected: true,
			limiter:  limiter,
		},
		{
			name:     "allow request",
			key:      "ip_not_allow",
			try:      10000,
			waitFor:  time.Microsecond,
			expected: false,
			limiter:  limiter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var result bool
			for i := int64(0); i < tt.try; i++ {
				result = tt.limiter.Allow(tt.key)
				time.Sleep(tt.waitFor)
			}

			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
