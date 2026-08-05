# cron-scheduler

Single-node background job scheduling pattern based on `robfig/cron/v3`.
Perfect for executing recurring tasks (e.g., daily emails, hourly database cleanups) in Go backends without needing complex distributed queues (like Celery or Asynq).

## Agent Usage Instructions

1. **Copy** `scheduler.go` into `<product-root>/<app-id>/internal/cron/`.
2. **Install deps**: 
   `go get github.com/robfig/cron/v3`
3. **Initialize and Start in `main.go`**:
   ```go
   import "your_app/internal/cron"

   scheduler := cron.NewScheduler()
   
   // Run every day at midnight
   scheduler.AddJob("0 0 * * *", func() {
       log.Println("Running daily cleanup...")
   })
   
   scheduler.Start()
   defer scheduler.Stop(context.Background())
   ```
