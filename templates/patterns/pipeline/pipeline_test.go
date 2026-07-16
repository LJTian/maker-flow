package pipeline

import "testing"

func TestFanOutAndMerge(t *testing.T) {
	in := make(chan int)
	go func() {
		for i := 1; i <= 4; i++ {
			in <- i
		}
		close(in)
	}()
	out := FanOut(2, in, func(v int) int { return v * 2 })
	sum := 0
	for v := range out {
		sum += v
	}
	if sum != 20 {
		t.Fatalf("sum=%d want 20", sum)
	}
}
