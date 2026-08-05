# storage-s3

S3-compatible storage pattern wrapper using `aws-sdk-go-v2`.
Highly optimized for configuring alternative object storage providers like Cloudflare R2 and MinIO without the bloat.

## Agent Usage Instructions

1. **Copy** `storage.go` into `<product-root>/<app-id>/internal/storage/`.
2. **Install deps**: 
   `go get github.com/aws/aws-sdk-go-v2 github.com/aws/aws-sdk-go-v2/config github.com/aws/aws-sdk-go-v2/credentials github.com/aws/aws-sdk-go-v2/service/s3`
3. **Initialize in `main.go`**:
   ```go
   import "your_app/internal/storage"

   s3Client, err := storage.NewClient(context.Background(), storage.Config{
       Endpoint:        os.Getenv("S3_ENDPOINT"), // e.g. https://<id>.r2.cloudflarestorage.com
       AccessKeyID:     os.Getenv("S3_ACCESS_KEY"),
       SecretAccessKey: os.Getenv("S3_SECRET_KEY"),
       Region:          os.Getenv("S3_REGION"), // "auto" for R2
       Bucket:          os.Getenv("S3_BUCKET"),
   })
   ```
4. **Usage**:
   ```go
   err := s3Client.UploadFile(ctx, "avatars/user_1.png", dataBytes, "image/png")
   ```
