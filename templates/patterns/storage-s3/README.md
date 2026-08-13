# storage-s3

**English** · [简体中文](README.zh-CN.md)

Decoupled object storage pattern: **`Storage` interface abstraction layer** with concrete implementations for **`R2StorageDriver`** (Cloudflare R2), **`MinIOStorageDriver`** (Docker / local MinIO), and **`MockStorageDriver`** (In-memory testing).

Tags: `storage` `s3` `r2` `minio` `upload` `docker`

## Decoupled Architecture

```
                    ┌───────────────────┐
                    │ Storage Interface │
                    └─────────┬─────────┘
                              │
          ┌───────────────────┼───────────────────┐
          ▼                   ▼                   ▼
  ┌───────────────┐   ┌───────────────┐   ┌───────────────┐
  │R2StorageDriver│   │MinIOStorage...│   │MockStorage... │
  │(Cloudflare R2)│   │(Docker MinIO) │   │  (In-Memory)  │
  └───────────────┘   └───────────────┘   └───────────────┘
```

## Usage

```go
import "your_app/internal/storage"

// Read config from env (STORAGE_MODE=r2 or minio)
cfg := storage.ConfigFromEnv()

// Factory constructor returns Storage interface
s, err := storage.NewStorage(ctx, cfg)
if err != nil {
    log.Fatal(err)
}

// Upload / Download via abstract Storage interface
err = s.UploadFile(ctx, "avatars/user_1.png", dataBytes, "image/png")
data, err := s.DownloadFile(ctx, "avatars/user_1.png")
```
