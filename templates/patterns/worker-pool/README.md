# worker-pool

**English** · [简体中文](README.zh-CN.md)

Fixed-size workers + buffered job channel + WaitGroup drain.

```go
p := pool.New(4, 32, fn)
p.Start(ctx)
_ = p.Submit(job)
<-ctx.Done()
p.Wait()
```

Tags: `concurrency` `pool` `channel`
