package metadata

import (
	"context"
	"encoding/json"

	"github.com/cdxy1/go-file-storage/internal/entity"
	"github.com/cdxy1/go-file-storage/internal/service"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type MetadataHandler struct {
	UnimplementedMetadataServiceServer
	svc *service.MetadataService
}

func NewMetadataHandler(svc *service.MetadataService) *MetadataHandler {
	return &MetadataHandler{svc: svc}
}

func (mh *MetadataHandler) GetAll(ctx context.Context) (*FilesMetadataResponse, error) {
	var sl []*FileMetadataResponse

	rows, err := mh.svc.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range rows {
		fmr := &FileMetadataResponse{
			Id:        v.Id,
			Name:      v.Name,
			Url:       v.Url,
			Size:      v.Size,
			Type:      v.Type,
			CreatedAt: v.CreatedAt.String(),
		}

		sl = append(sl, fmr)
	}

	return &FilesMetadataResponse{Files: sl}, nil
}

func (mh *MetadataHandler) GetById(ctx context.Context, req *FileMetadataRequest) (*FileMetadataResponse, error) {
	row, err := mh.svc.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &FileMetadataResponse{
		Id:        row.Id,
		Name:      row.Name,
		Url:       row.Url,
		Size:      row.Size,
		Type:      row.Type,
		CreatedAt: row.CreatedAt.String(),
	}, nil
}

func (fh *MetadataHandler) HandleMessage(msg []byte, offset kafka.Offset) error {
	data := entity.Metadata{}
	if err := json.Unmarshal(msg, &data); err != nil {
		return err
	}

	return nil
}
