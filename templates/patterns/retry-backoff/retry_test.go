package retry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestDoSucceedsAfterRetries(t *testing.T) {
	n := 0
	err := Do(context.Background(), Options{MaxAttempts: 3, Initial: time.Millisecond}, func(ctx context.Context) error {
		n++
		if n < 3 {
			return errors.New("fail")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Fatalf("n=%d", n)
	}
}
