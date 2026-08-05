# storage-s3

使用 `aws-sdk-go-v2` 封装的兼容 S3 的对象存储模式。
专为接入非传统云存储（如 Cloudflare R2、MinIO 等高性价比方案）做了路径与端点解析器的深度适配，即插即用。

## Agent 使用说明

1. **复制** `storage.go` 到 `<产品根>/<app-id>/internal/storage/`。
2. **安装依赖**: 
   `go get github.com/aws/aws-sdk-go-v2 github.com/aws/aws-sdk-go-v2/config github.com/aws/aws-sdk-go-v2/credentials github.com/aws/aws-sdk-go-v2/service/s3`
3. **在 `main.go` 中初始化**:
   ```go
   import "your_app/internal/storage"

   s3Client, err := storage.NewClient(context.Background(), storage.Config{
       Endpoint:        os.Getenv("S3_ENDPOINT"), // 例如 Cloudflare R2 的 API
       AccessKeyID:     os.Getenv("S3_ACCESS_KEY"),
       SecretAccessKey: os.Getenv("S3_SECRET_KEY"),
       Region:          os.Getenv("S3_REGION"), // R2 一般使用 "auto"
       Bucket:          os.Getenv("S3_BUCKET"),
   })
   ```
4. **调用示例**:
   ```go
   err := s3Client.UploadFile(ctx, "avatars/user_1.png", dataBytes, "image/png")
   ```
