package cron

import (
	"context"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	scheduler := NewScheduler()

	_, err := scheduler.AddJob("* * * * *", func() {})
	if err != nil {
		t.Fatalf("Failed to add job: %v", err)
	}

	scheduler.Start()

	// Keep unit tests fast: only exercise Start + graceful Stop (no minute wait).
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = scheduler.Stop(ctx)
	if err != nil {
		t.Fatalf("Failed to stop scheduler: %v", err)
	}
}
