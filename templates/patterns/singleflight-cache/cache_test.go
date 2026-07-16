package cache

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestSingleflight(t *testing.T) {
	c := New(time.Second)
	var calls atomic.Int64
	load := func() (any, error) {
		calls.Add(1)
		time.Sleep(20 * time.Millisecond)
		return "ok", nil
	}
	done := make(chan struct{}, 5)
	for i := 0; i < 5; i++ {
		go func() {
			_, err, _ := c.GetOrLoad("k", load)
			if err != nil {
				t.Errorf("err %v", err)
			}
			done <- struct{}{}
		}()
	}
	for i := 0; i < 5; i++ {
		<-done
	}
	if calls.Load() != 1 {
		t.Fatalf("calls=%d want 1", calls.Load())
	}
}
