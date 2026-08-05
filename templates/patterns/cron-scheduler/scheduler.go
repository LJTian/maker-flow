package cron

import (
	"context"
	"log"

	"github.com/robfig/cron/v3"
)

// Scheduler wraps the robfig/cron library
type Scheduler struct {
	cron *cron.Cron
}

// NewScheduler initializes a new cron scheduler with standard parser (minute, hour, dom, month, dow)
func NewScheduler() *Scheduler {
	c := cron.New(cron.WithParser(cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
	)))
	return &Scheduler{
		cron: c,
	}
}

// AddJob registers a function to run on a standard cron schedule
func (s *Scheduler) AddJob(schedule string, job func()) (cron.EntryID, error) {
	return s.cron.AddFunc(schedule, job)
}

// Start begins executing the registered cron jobs in a background goroutine
func (s *Scheduler) Start() {
	s.cron.Start()
	log.Println("Cron scheduler started")
}

// Stop gracefully stops the cron scheduler and waits for running jobs to finish
func (s *Scheduler) Stop(ctx context.Context) error {
	log.Println("Stopping cron scheduler...")
	stopCtx := s.cron.Stop()
	
	select {
	case <-stopCtx.Done():
		log.Println("Cron scheduler stopped gracefully")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
