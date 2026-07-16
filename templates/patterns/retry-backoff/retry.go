package retry

import (
	"context"
	"time"
)

type Options struct {
	MaxAttempts int
	Initial     time.Duration
	MaxDelay    time.Duration
	Multiplier  float64
}

func Do(ctx context.Context, opt Options, fn func(context.Context) error) error {
	if opt.MaxAttempts < 1 {
		opt.MaxAttempts = 1
	}
	if opt.Initial <= 0 {
		opt.Initial = 100 * time.Millisecond
	}
	if opt.MaxDelay <= 0 {
		opt.MaxDelay = 5 * time.Second
	}
	if opt.Multiplier < 1 {
		opt.Multiplier = 2
	}

	delay := opt.Initial
	var err error
	for attempt := 1; attempt <= opt.MaxAttempts; attempt++ {
		if err = fn(ctx); err == nil {
			return nil
		}
		if attempt == opt.MaxAttempts {
			break
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
		}
		delay = time.Duration(float64(delay) * opt.Multiplier)
		if delay > opt.MaxDelay {
			delay = opt.MaxDelay
		}
	}
	return err
}
