package minio

import (
	"context"
	"log/slog"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/cdxy1/go-file-storage/internal/config"
)

type Minio struct {
	MinioClient *minio.Client
	BucketName  string
	Logger      *slog.Logger
}

func NewMinio(logger *slog.Logger) (*Minio, error) {
	cfg := config.GetConfig()

	logger.Info("Initializing MinIO client",
		"endpoint", cfg.Minio.Endpoint,
		"useSSL", cfg.Minio.UseSSL,
	)

	mc, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.User, cfg.Minio.Password, ""),
		Secure: cfg.Minio.UseSSL,
	})
	if err != nil {
		logger.Error("Failed to initialize MinIO client",
			"endpoint", cfg.Minio.Endpoint,
			"error", err,
		)
		return nil, err
	}

	if err := CreateBucket(mc, cfg.Minio.Bucket, logger); err != nil {
		logger.Error("Bucket creation failed",
			"bucket", cfg.Minio.Bucket,
			"error", err,
		)
	}

	logger.Info("MinIO client initialized successfully",
		"bucket", cfg.Minio.Bucket,
	)

	minio := &Minio{
		MinioClient: mc,
		BucketName:  cfg.Minio.Bucket,
		Logger:      logger,
	}

	return minio, nil
}

func CreateBucket(client *minio.Client, bucketName string, logger *slog.Logger) error {
	ctx := context.Background()

	exist, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		logger.Error("Failed to check bucket existence",
			"bucket", bucketName,
			"error", err,
		)
		return err
	}

	if !exist {
		logger.Info("Bucket not found, creating new one", "bucket", bucketName)
		if err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
			logger.Error("Failed to create bucket",
				"bucket", bucketName,
				"error", err,
			)
			return err
		}
		logger.Info("Bucket successfully created", "bucket", bucketName)
	} else {
		logger.Debug("Bucket already exists", "bucket", bucketName)
	}

	return nil
}
