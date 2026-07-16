package pipeline

import "sync"

// FanOut runs n workers reading from in and writing to a shared out.
func FanOut[T any](n int, in <-chan T, fn func(T) T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- fn(v)
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
func Merge[T any](chans ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(chans))
	for _, ch := range chans {
		ch := ch
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
