package pool

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestPoolProcessesJobs(t *testing.T) {
	var n atomic.Int64
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := New(2, 8, func(ctx context.Context, job int) {
		n.Add(1)
	})
	p.Start(ctx)
	for i := 0; i < 5; i++ {
		if !p.Submit(i) {
			t.Fatal("submit failed")
		}
	}
	time.Sleep(50 * time.Millisecond)
	cancel()
	p.Wait()
	if n.Load() != 5 {
		t.Fatalf("got %d want 5", n.Load())
	}
}
