package repo

import (
	"context"
	"io"
	"log/slog"

	ms "github.com/cdxy1/minio-go/internal/storage/minio"
	"github.com/minio/minio-go/v7"
)

type File struct {
	mc *ms.Minio
}

func NewFileRepo(logger *slog.Logger) (*File, error) {
	ms, err := ms.NewMinio(logger)
	if err != nil {
		return nil, err
	}

	return &File{ms}, nil
}

func (f *File) Put(ctx context.Context, objName string, reader io.Reader, size int64) error {
	_, err := f.mc.MinioClient.PutObject(ctx, f.mc.BucketName, objName, reader, size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}

	return nil
}

func (f *File) GetByName(ctx context.Context, objName string) (*minio.Object, error) {
	obj, err := f.mc.MinioClient.GetObject(ctx, f.mc.BucketName, objName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return obj, nil
}
