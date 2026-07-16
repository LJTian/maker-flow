[English](README.md) · **简体中文**

# worker-pool

固定数量 worker + 有缓冲 job channel + WaitGroup 排空。

```go
p := pool.New(4, 32, fn)
p.Start(ctx)
_ = p.Submit(job)
<-ctx.Done()
p.Wait()
```

Tags: `concurrency` `pool` `channel`
