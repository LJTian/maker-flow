package worker

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

type Job struct {
	ID      int
	Payload string
}

type Pool struct {
	workers int
	jobs    chan Job
	wg      sync.WaitGroup
	logger  *slog.Logger
}

func NewPool(workers, queueSize int, logger *slog.Logger) *Pool {
	if workers < 1 {
		workers = 1
	}
	if queueSize < 1 {
		queueSize = 1
	}
	return &Pool{
		workers: workers,
		jobs:    make(chan Job, queueSize),
		logger:  logger,
	}
}

func (p *Pool) Start(ctx context.Context) {
	for i := 0; i < p.workers; i++ {
		id := i
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			p.loop(ctx, id)
		}()
	}
	go func() {
		<-ctx.Done()
		close(p.jobs)
	}()
}

func (p *Pool) Submit(job Job) error {
	select {
	case p.jobs <- job:
		return nil
	default:
		p.logger.Warn("queue full, drop", "job_id", job.ID)
		return ErrQueueFull
	}
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) loop(ctx context.Context, workerID int) {
	for job := range p.jobs {
		p.process(ctx, workerID, job)
	}
	p.logger.Info("worker stopped", "worker", workerID)
}

func (p *Pool) process(ctx context.Context, workerID int, job Job) {
	select {
	case <-ctx.Done():
		return
	default:
	}
	// demo work — replace with business logic
	time.Sleep(100 * time.Millisecond)
	p.logger.Info("job done", "worker", workerID, "job_id", job.ID, "payload", job.Payload)
}
