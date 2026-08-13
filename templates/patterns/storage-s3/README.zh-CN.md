# storage-s3

[English](README.md) · **简体中文**

解耦的对象存储模式组件：提供统一的 **`Storage` 接口抽象层**，由 **`R2StorageDriver`**（Cloudflare R2）、**`MinIOStorageDriver`**（Docker / 本地 MinIO）和 **`MockStorageDriver`**（内存测试驱动）提供具体实现。

标签: `storage` `s3` `r2` `minio` `upload` `docker`

## 解耦架构设计

```
                    ┌───────────────────┐
                    │ Storage 接口抽象层 │
                    └─────────┬─────────┘
                              │
          ┌───────────────────┼───────────────────┐
          ▼                   ▼                   ▼
  ┌───────────────┐   ┌───────────────┐   ┌───────────────┐
  │R2StorageDriver│   │MinIOStorage...│   │MockStorage... │
  │(Cloudflare R2)│   │(Docker MinIO) │   │  (内存测试)   │
  └───────────────┘   └───────────────┘   └───────────────┘
```

## 使用示例

```go
import "your_app/internal/storage"

// 从环境变量读取配置 (STORAGE_MODE=r2 或 minio)
cfg := storage.ConfigFromEnv()

// 工厂构造函数返回 Storage 接口对象
s, err := storage.NewStorage(ctx, cfg)
if err != nil {
    log.Fatal(err)
}

// 通过统一的 Storage 接口执行文件上传/下载
err = s.UploadFile(ctx, "avatars/user_1.png", dataBytes, "image/png")
data, err := s.DownloadFile(ctx, "avatars/user_1.png")
```
