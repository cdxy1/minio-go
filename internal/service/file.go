package service

import (
	"bytes"
	"context"
	"io"

	"github.com/cdxy1/minio-go/internal/repo"
	"log/slog"
)

type FileService struct {
	Repo   *repo.File
	Logger *slog.Logger
}

func NewFileService(repo *repo.File, logger *slog.Logger) *FileService {
	return &FileService{Repo: repo, Logger: logger}
}

func (fs *FileService) DownloadFile(ctx context.Context, objName string) ([]byte, error) {
	obj, err := fs.Repo.GetByName(ctx, objName)
	if err != nil {
		fs.Logger.Error("File not found", "file", objName, "error", err)
		return nil, err
	}
	defer obj.Close()

	data, err := io.ReadAll(obj)
	if err != nil {
		fs.Logger.Error("Unable to read file", "file", objName, "error", err)
		return nil, err
	}

	fs.Logger.Info("File successfully read", "file", objName, "size", len(data))
	return data, nil
}

func (fs *FileService) UploadFile(ctx context.Context, objName string, objData []byte) (string, error) {
	rdr := bytes.NewReader(objData)
	dl := int64(rdr.Len())

	if err := fs.Repo.Put(ctx, objName, rdr, dl); err != nil {
		fs.Logger.Error("Error while reading", "file", objName)
		return "", err
	}

	fs.Logger.Info("Successfully uploaded", "file", objName)
	return objName, nil
}
