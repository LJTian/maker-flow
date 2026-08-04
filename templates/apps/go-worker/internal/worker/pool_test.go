package worker

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	p := NewPool(2, 5, logger)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	p.Start(ctx)

	for i := 1; i <= 3; i++ {
		err := p.Submit(Job{ID: i, Payload: "test"})
		if err != nil {
			t.Fatalf("failed to submit job: %v", err)
		}
	}

	cancel() // trigger shutdown
	p.Wait() // ensure clean shutdown doesn't block forever
}
