# storage-s3

[English](README.md) · **简体中文**

基于 `aws-sdk-go-v2` 封装的 S3 兼容对象存储组件。针对线上 **Cloudflare R2** 与本地 **Docker MinIO** 做了专门优化。

标签: `storage` `s3` `r2` `minio` `upload` `docker`

## 部署策略

- **本地开发**：Docker MinIO（使用 `compose.snippet.yml`，端点 `http://localhost:9000`）。
- **线上生产 (Cloudflare R2)**：
  - `S3_ENDPOINT=https://<account_id>.r2.cloudflarestorage.com`
  - `S3_REGION=auto`
  - `S3_ACCESS_KEY=<r2_access_key_id>`
  - `S3_SECRET_KEY=<r2_secret_access_key>`

## 免费额度说明

- **Cloudflare R2 免费额度**：10 GB 存储空间/月，1000 万次读取/月，100 万次写入/月，**0 元出站流量费**。

## Agent 使用说明

1. **复制** `storage.go` 到 `<product-root>/<app-id>/internal/storage/`。
2. **安装依赖**：
   `go get github.com/aws/aws-sdk-go-v2 github.com/aws/aws-sdk-go-v2/config github.com/aws/aws-sdk-go-v2/credentials github.com/aws/aws-sdk-go-v2/service/s3`
3. **在 `main.go` 中初始化**：
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
