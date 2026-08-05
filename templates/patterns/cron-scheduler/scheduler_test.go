package cron

import (
	"context"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	scheduler := NewScheduler()
	
	executed := false
	_, err := scheduler.AddJob("* * * * *", func() {
		executed = true
	})
	
	if err != nil {
		t.Fatalf("Failed to add job: %v", err)
	}

	scheduler.Start()
	
	// We won't actually wait a full minute to test it to keep unit tests fast, 
	// we just test start and graceful stop.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = scheduler.Stop(ctx)
	if err != nil {
		t.Fatalf("Failed to stop scheduler: %v", err)
	}
}
