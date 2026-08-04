package pipeline

import (
	"context"
	"sync"
)

// FanOut runs n workers reading from in and writing to a shared out.
func FanOut[T any](ctx context.Context, n int, in <-chan T, fn func(T) T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					res := fn(v)
					select {
					case out <- res:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Merge fans-in multiple channels into one.
func Merge[T any](ctx context.Context, chans ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(chans))
	for _, ch := range chans {
		ch := ch
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok {
						return
					}
					select {
					case out <- v:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
