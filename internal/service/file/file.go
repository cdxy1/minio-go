package file

import (
	"context"

	"github.com/cdxy1/go-file-storage/internal/repo"
	"github.com/cdxy1/go-file-storage/internal/entity"
)

type FileService struct {
	repo *repo.File
}

func (fs *FileService) GetFile(ctx context.Context, id int) (*entity.File, error) {
	f, err := fs.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}
	
	return f, nil
}
