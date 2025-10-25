package service

import (
	"context"

	"github.com/cdxy1/go-file-storage/internal/entity"
	"github.com/cdxy1/go-file-storage/internal/repo"
)

type MetadataService struct {
	Repo *repo.Metadata
}

func NewMetadataService(repo *repo.Metadata) *MetadataService {
	return &MetadataService{Repo: repo}
}

func (ms *MetadataService) GetById(ctx context.Context, id string) (*entity.Metadata, error) {
	f, err := ms.Repo.GetByID(ctx, id)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return f, nil
}

func (ms *MetadataService) CreateFile(ctx context.Context, u *entity.Metadata) error {
	if err := ms.Repo.Create(ctx, u); err != nil {
		return err
	}
	return nil
}

func (ms *MetadataService) GetAll(ctx context.Context) ([]entity.Metadata, error) {
	sl, err := ms.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return sl, nil
}
