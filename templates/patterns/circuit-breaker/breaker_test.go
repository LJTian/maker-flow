package breaker

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestTripsOpen(t *testing.T) {
	b := New(2, 50*time.Millisecond)
	_ = b.Do(func() error { return errors.New("x") })
	_ = b.Do(func() error { return errors.New("x") })
	if b.State() != Open {
		t.Fatalf("state=%v want Open", b.State())
	}
	if err := b.Do(func() error { return nil }); !errors.Is(err, ErrOpen) {
		t.Fatalf("err=%v", err)
	}
	time.Sleep(60 * time.Millisecond)
	if b.State() != HalfOpen {
		t.Fatalf("state=%v want HalfOpen", b.State())
	}
	if err := b.Do(func() error { return nil }); err != nil {
		t.Fatal(err)
	}
	if b.State() != Closed {
		t.Fatalf("state=%v want Closed", b.State())
	}
}

func TestHalfOpenConcurrency(t *testing.T) {
	b := New(1, 50*time.Millisecond)
	_ = b.Do(func() error { return errors.New("x") })
	time.Sleep(60 * time.Millisecond)

	var wg sync.WaitGroup
	var succ atomic.Int32
	var rej atomic.Int32

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := b.Do(func() error {
				time.Sleep(10 * time.Millisecond) // Ensure it holds the halfOpen slot
				return nil
			})
			if err == nil {
				succ.Add(1)
			} else if errors.Is(err, ErrOpen) {
				rej.Add(1)
			}
		}()
	}
	wg.Wait()
	
	// Because the first success transitions state to Closed, subsequent calls 
	// in the loop might hit Closed and succeed if they start after the 10ms sleep.
	// But any calls made DURING the 10ms sleep should be rejected.
	// We at least expect 1 success and some rejections.
	if succ.Load() < 1 {
		t.Errorf("succ=%d want >= 1", succ.Load())
	}
	if rej.Load() == 0 {
		t.Errorf("rej=%d want > 0", rej.Load())
	}
}
