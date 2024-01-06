package limiter_test

import (
	"testing"

	"github.com/k-vanio/rage-limiter/internal/core/limiter"
)

func TestAllow(t *testing.T) {
	l := limiter.New()

	if !l.Allow("test") {
		t.Error("expected true")
	}
}
