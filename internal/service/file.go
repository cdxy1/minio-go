package service

import (
	"bytes"
	"context"
	"io"

	"github.com/cdxy1/go-file-storage/internal/repo"
)

type FileService struct {
	Repo *repo.File
}

func NewFileService(repo *repo.File) *FileService {
	return &FileService{Repo: repo}
}

func (fs *FileService) DownloadFile(ctx context.Context, objName string) ([]byte, error) {
	obj, err := fs.Repo.GetByName(ctx, objName)

	if err != nil {
		return nil, err
	}
	defer obj.Close()

	data, err := io.ReadAll(obj)

	if err != nil {
		println(obj, err.Error())
		return nil, err
	}

	return data, nil
}

func (fs *FileService) UploadFile(ctx context.Context, objName string, objData []byte) (string, error) {
	rdr := bytes.NewReader(objData)
	dl := int64(rdr.Len())

	if err := fs.Repo.Put(ctx, objName, rdr, dl); err != nil {
		return "", err
	}
	return objName, nil
}
