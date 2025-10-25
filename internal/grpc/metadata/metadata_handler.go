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

	return &FilesMetadataResponse{Files: []*FileMetadataResponse{}}, nil
}

func (mh *MetadataHandler) GetById(ctx context.Context, req *FileMetadataRequest) (*FileMetadataResponse, error) {

	return  &FileMetadataResponse{}, nil
}

func (fh *MetadataHandler) HandleMessage(msg []byte, offset kafka.Offset) error {
	data := entity.Metadata{}
	if err := json.Unmarshal(msg, &data); err != nil {
		return err
	}

	return nil
}
