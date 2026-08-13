package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Storage is the abstract object storage interface.
type Storage interface {
	UploadFile(ctx context.Context, key string, data []byte, contentType string) error
	DownloadFile(ctx context.Context, key string) ([]byte, error)
}

// Config defines settings for storage driver initialization.
type Config struct {
	Mode            string // "r2", "minio", or "mock"
	Endpoint        string // Endpoint URL
	AccessKeyID     string
	SecretAccessKey string
	Region          string // "auto" for R2, "us-east-1" for MinIO
	Bucket          string
}

// ConfigFromEnv reads storage configuration from environment variables.
func ConfigFromEnv() Config {
	mode := os.Getenv("STORAGE_MODE")
	if mode == "" {
		mode = "minio"
	}
	endpoint := os.Getenv("S3_ENDPOINT")
	if endpoint == "" && mode == "minio" {
		endpoint = "http://localhost:9000"
	}
	region := os.Getenv("S3_REGION")
	if region == "" {
		if mode == "r2" {
			region = "auto"
		} else {
			region = "us-east-1"
		}
	}

	return Config{
		Mode:            mode,
		Endpoint:        endpoint,
		AccessKeyID:     os.Getenv("S3_ACCESS_KEY"),
		SecretAccessKey: os.Getenv("S3_SECRET_KEY"),
		Region:          region,
		Bucket:          os.Getenv("S3_BUCKET"),
	}
}

// NewStorage is the factory constructor returning the Storage interface.
func NewStorage(ctx context.Context, cfg Config) (Storage, error) {
	switch cfg.Mode {
	case "r2":
		return NewR2StorageDriver(ctx, cfg)
	case "minio":
		return NewMinIOStorageDriver(ctx, cfg)
	case "mock":
		return NewMockStorageDriver(), nil
	default:
		return NewMinIOStorageDriver(ctx, cfg)
	}
}

// --- R2StorageDriver (Cloudflare R2) ---

type R2StorageDriver struct {
	s3     *s3.Client
	bucket string
}

func NewR2StorageDriver(ctx context.Context, cfg Config) (*R2StorageDriver, error) {
	if cfg.Endpoint == "" || cfg.AccessKeyID == "" || cfg.SecretAccessKey == "" {
		return nil, fmt.Errorf("r2 mode requires S3_ENDPOINT, S3_ACCESS_KEY, and S3_SECRET_KEY")
	}

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           cfg.Endpoint,
			SigningRegion: cfg.Region,
		}, nil
	})

	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, "")),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsCfg)

	return &R2StorageDriver{
		s3:     client,
		bucket: cfg.Bucket,
	}, nil
}

func (r *R2StorageDriver) UploadFile(ctx context.Context, key string, data []byte, contentType string) error {
	_, err := r.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
	})
	return err
}

func (r *R2StorageDriver) DownloadFile(ctx context.Context, key string) ([]byte, error) {
	out, err := r.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer out.Body.Close()
	return io.ReadAll(out.Body)
}

// --- MinIOStorageDriver (Local Docker MinIO) ---

type MinIOStorageDriver struct {
	s3     *s3.Client
	bucket string
}

func NewMinIOStorageDriver(ctx context.Context, cfg Config) (*MinIOStorageDriver, error) {
	endpoint := cfg.Endpoint
	if endpoint == "" {
		endpoint = "http://localhost:9000"
	}
	accessKey := cfg.AccessKeyID
	if accessKey == "" {
		accessKey = "minioadmin"
	}
	secretKey := cfg.SecretAccessKey
	if secretKey == "" {
		secretKey = "minioadmin"
	}
	region := cfg.Region
	if region == "" {
		region = "us-east-1"
	}

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           endpoint,
			SigningRegion: region,
		}, nil
	})

	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return &MinIOStorageDriver{
		s3:     client,
		bucket: cfg.Bucket,
	}, nil
}

func (m *MinIOStorageDriver) UploadFile(ctx context.Context, key string, data []byte, contentType string) error {
	_, err := m.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(m.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
	})
	return err
}

func (m *MinIOStorageDriver) DownloadFile(ctx context.Context, key string) ([]byte, error) {
	out, err := m.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(m.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer out.Body.Close()
	return io.ReadAll(out.Body)
}

// --- MockStorageDriver (In-Memory for unit testing) ---

type MockStorageDriver struct {
	mu    sync.RWMutex
	files map[string][]byte
}

func NewMockStorageDriver() *MockStorageDriver {
	return &MockStorageDriver{
		files: make(map[string][]byte),
	}
}

func (m *MockStorageDriver) UploadFile(ctx context.Context, key string, data []byte, contentType string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.files[key] = data
	return nil
}

func (m *MockStorageDriver) DownloadFile(ctx context.Context, key string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data, exists := m.files[key]
	if !exists {
		return nil, fmt.Errorf("file not found: %s", key)
	}
	return data, nil
}
