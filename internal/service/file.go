package service

import (
	"context"

	"github.com/cdxy1/go-file-storage/internal/entity"
	"github.com/cdxy1/go-file-storage/internal/repo"
)

type FileService struct {
	Repo *repo.File
}

func (fs *FileService) GetFile(ctx context.Context, id int) (*entity.File, error) {
	f, err := fs.Repo.GetByID(ctx, id)

	if err != nil {
		println(err.Error())
		return nil, err
	}

	return f, nil
}

func (fs *FileService) CreateFile(ctx context.Context, u *entity.File) error {
	if err := fs.Repo.Create(ctx, u); err != nil {
		return err
	}
	return nil
}
