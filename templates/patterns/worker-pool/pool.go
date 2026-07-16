package pool

import (
	"context"
	"sync"
)

type Handler[T any] func(ctx context.Context, job T)

type Pool[T any] struct {
	workers int
	jobs    chan T
	handle  Handler[T]
	wg      sync.WaitGroup
}

func New[T any](workers, queue int, handle Handler[T]) *Pool[T] {
	if workers < 1 {
		workers = 1
	}
	if queue < 1 {
		queue = 1
	}
	return &Pool[T]{
		workers: workers,
		jobs:    make(chan T, queue),
		handle:  handle,
	}
}

func (p *Pool[T]) Start(ctx context.Context) {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for job := range p.jobs {
				p.handle(ctx, job)
			}
		}()
	}
	go func() {
		<-ctx.Done()
		close(p.jobs)
	}()
}

func (p *Pool[T]) Submit(job T) bool {
	select {
	case p.jobs <- job:
		return true
	default:
		return false
	}
}

func (p *Pool[T]) Wait() { p.wg.Wait() }
