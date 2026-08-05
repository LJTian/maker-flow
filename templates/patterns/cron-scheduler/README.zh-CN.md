# cron-scheduler

基于 `robfig/cron/v3` 的单机定时任务调度模式。
非常适合在早期的 Go 后端中执行周期性任务（例如：发送每日数据报表、每小时清理数据库垃圾、检查订阅过期），而不需要引入复杂的分布式队列（如 Asynq 或 Celery）。

## Agent 使用说明

1. **复制** `scheduler.go` 到 `<产品根>/<app-id>/internal/cron/`。
2. **安装依赖**: 
   `go get github.com/robfig/cron/v3`
3. **在 `main.go` 中初始化与启动**:
   ```go
   import "your_app/internal/cron"

   scheduler := cron.NewScheduler()
   
   // 每天午夜 00:00 运行
   scheduler.AddJob("0 0 * * *", func() {
       log.Println("运行每日清理任务...")
   })
   
   scheduler.Start()
   
   // 配合系统的优雅停机处理退出
   defer scheduler.Stop(context.Background())
   ```
