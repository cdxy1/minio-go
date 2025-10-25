package service

import (
	"context"

	"github.com/cdxy1/go-file-storage/internal/entity"
	"github.com/cdxy1/go-file-storage/internal/repo"
)

type MetadataService struct {
	Repo *repo.Metadata
}

func (fs *MetadataService) GetFile(ctx context.Context, id string) (*entity.Metadata, error) {
	f, err := fs.Repo.GetByID(ctx, id)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return f, nil
}

func (fs *MetadataService) CreateFile(ctx context.Context, u *entity.Metadata) error {
	if err := fs.Repo.Create(ctx, u); err != nil {
		return err
	}
	return nil
}
