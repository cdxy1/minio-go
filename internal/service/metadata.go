package service

import (
	"context"
	"log/slog"

	"github.com/cdxy1/go-file-storage/internal/entity"
	"github.com/cdxy1/go-file-storage/internal/repo"
)

type MetadataService struct {
	Repo   *repo.Metadata
	Logger *slog.Logger
}

func NewMetadataService(repo *repo.Metadata, logger *slog.Logger) *MetadataService {
	return &MetadataService{
		Repo:   repo,
		Logger: logger,
	}
}

func (ms *MetadataService) GetById(ctx context.Context, id string) (*entity.Metadata, error) {
	ms.Logger.Info("GetById called", "id", id)

	f, err := ms.Repo.GetByID(ctx, id)
	if err != nil {
		ms.Logger.Error("GetById failed", "id", id, "err", err)
		return nil, err
	}

	ms.Logger.Info("GetById completed", "id", id)
	return f, nil
}

func (ms *MetadataService) CreateFile(ctx context.Context, u *entity.Metadata) error {
	if err := ms.Repo.Create(ctx, u); err != nil {
		ms.Logger.Error("CreateFile failed", "file", u.Name, "err", err)
		return err
	}

	ms.Logger.Info("CreateFile completed", "file", u.Name)
	return nil
}

func (ms *MetadataService) GetAll(ctx context.Context) ([]entity.Metadata, error) {
	ms.Logger.Info("GetAll called")

	sl, err := ms.Repo.GetAll(ctx)
	if err != nil {
		ms.Logger.Error("GetAll failed", "err", err)
		return nil, err
	}

	ms.Logger.Info("GetAll completed", "count", len(sl))
	return sl, nil
}
