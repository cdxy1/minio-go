package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/cdxy1/go-file-storage/internal/config"
)

func NewMinio() (*minio.Client, error) {
	cfg := config.GetConfig()

	minioClient, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.User, cfg.Minio.Password, ""),
		Secure: cfg.Minio.UseSSL,
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
