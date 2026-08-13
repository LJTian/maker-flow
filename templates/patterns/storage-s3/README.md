# storage-s3

**English** · [简体中文](README.zh-CN.md)

S3-compatible object storage wrapper using `aws-sdk-go-v2`. Optimized for **Cloudflare R2** in production and **MinIO** in local development.

Tags: `storage` `s3` `r2` `minio` `upload` `docker`

## Deployment Strategy

- **Local Development**: Docker MinIO (`S3_ENDPOINT=http://localhost:9000`, `S3_REGION=us-east-1`, `S3_ACCESS_KEY=minioadmin`, `S3_SECRET_KEY=minioadmin`).
- **Online Production (Cloudflare R2)**:
  - `S3_ENDPOINT=https://<account_id>.r2.cloudflarestorage.com`
  - `S3_REGION=auto`
  - `S3_ACCESS_KEY=<r2_access_key_id>`
  - `S3_SECRET_KEY=<r2_secret_access_key>`

## Free Tier & Cost

- **Cloudflare R2**: 10 GB storage / month, 10M reads / month, 1M writes / month, **$0.00 egress bandwidth fees**.

## Agent Usage Instructions

1. **Copy** `storage.go` into `<product-root>/<app-id>/internal/storage/`.
2. **Install deps**: 
   `go get github.com/aws/aws-sdk-go-v2 github.com/aws/aws-sdk-go-v2/config github.com/aws/aws-sdk-go-v2/credentials github.com/aws/aws-sdk-go-v2/service/s3`
3. **Initialize in `main.go`**:
   ```go
   import "your_app/internal/storage"

   s3Client, err := storage.NewClient(context.Background(), storage.Config{
       Endpoint:        os.Getenv("S3_ENDPOINT"),
       AccessKeyID:     os.Getenv("S3_ACCESS_KEY"),
       SecretAccessKey: os.Getenv("S3_SECRET_KEY"),
       Region:          os.Getenv("S3_REGION"),
       Bucket:          os.Getenv("S3_BUCKET"),
   })
   ```
