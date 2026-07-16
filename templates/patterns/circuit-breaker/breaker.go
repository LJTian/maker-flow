package breaker

import (
	"errors"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

var ErrOpen = errors.New("circuit open")

type Breaker struct {
	mu               sync.Mutex
	state            State
	failures         int
	FailureThreshold int
	OpenTimeout      time.Duration
	openedAt         time.Time
}

func New(failureThreshold int, openTimeout time.Duration) *Breaker {
	if failureThreshold < 1 {
		failureThreshold = 1
	}
	if openTimeout <= 0 {
		openTimeout = 5 * time.Second
	}
	return &Breaker{
		state:            Closed,
		FailureThreshold: failureThreshold,
		OpenTimeout:      openTimeout,
	}
}

func (b *Breaker) State() State {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.maybeHalfOpenLocked()
	return b.state
}

func (b *Breaker) Do(fn func() error) error {
	b.mu.Lock()
	b.maybeHalfOpenLocked()
	if b.state == Open {
		b.mu.Unlock()
		return ErrOpen
	}
	b.mu.Unlock()

	err := fn()

	b.mu.Lock()
	defer b.mu.Unlock()
	if err != nil {
		b.failures++
		if b.failures >= b.FailureThreshold {
			b.state = Open
			b.openedAt = time.Now()
		}
		return err
	}
	b.failures = 0
	b.state = Closed
	return nil
}

func (b *Breaker) maybeHalfOpenLocked() {
	if b.state == Open && time.Since(b.openedAt) >= b.OpenTimeout {
		b.state = HalfOpen
	}
}
