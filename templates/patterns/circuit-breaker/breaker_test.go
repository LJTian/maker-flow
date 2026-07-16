package breaker

import (
	"errors"
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
