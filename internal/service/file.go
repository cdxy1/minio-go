package service

import (
	"bytes"
	"context"

	"github.com/cdxy1/go-file-storage/internal/repo"
	"github.com/minio/minio-go/v7"
)

type FileService struct {
	repo *repo.File
}

func (fs *FileService) DownloadFile(ctx context.Context, objName string) (*minio.Object, error){
	obj, err := fs.repo.GetByName(ctx, objName)

	if err != nil {
		return nil, err
	}

	return obj, err
}

func (fs *FileService) UploadFile(ctx context.Context, objName string, objData []byte) error {
	rdr := bytes.NewReader(objData)
	dl := int64(rdr.Len())

	if err := fs.repo.Put(ctx, objName, rdr, dl); err != nil {
		return err
	}
	return nil
}
