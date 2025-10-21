package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/cdxy1/go-file-storage/internal/config"
)

type Minio struct {
	MinioClient *minio.Client
	BucketName  string
}

func NewMinio() (*Minio, error) {
	cfg := config.GetConfig()

	mc, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.User, cfg.Minio.Password, ""),
		Secure: cfg.Minio.UseSSL,
	})

	if err != nil {
		println(err.Error())
		return nil, err
	}

	if err := CreateBucket(mc, cfg.Minio.Bucket); err != nil {
		return nil, err
	}

	minio := &Minio{
		MinioClient: mc,
		BucketName:  cfg.Minio.Bucket,
	}

	return minio, nil
}

func CreateBucket(client *minio.Client, bucketName string) error {
	exist, _ := client.BucketExists(context.Background(), bucketName)

	if !exist {
		if err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{}); err != nil {
			return err
		}
	}
	return nil
}
