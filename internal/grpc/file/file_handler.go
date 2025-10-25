package file

import (
	"context"

	"github.com/cdxy1/go-file-storage/internal/lib"
	"github.com/cdxy1/go-file-storage/internal/service"
	"github.com/cdxy1/go-file-storage/internal/infra/kafka/producer"
)

type FileHandler struct {
	UnimplementedFileServiceServer
	svc *service.FileService
	kafkaProducer *producer.Producer
}

func NewFileHandler(svc *service.FileService) *FileHandler {
	return &FileHandler{svc: svc}
}

func (fh *FileHandler) UploadFile(ctx context.Context, req *UploadFileRequest) (*UploadFileResponse, error) {
	name, err := fh.svc.UploadFile(ctx, req.Name, req.Data)
	if err != nil {
		return nil, err
	}

	msg, err := lib.ExtractMetadata(req.Name, req.Data)
	if err != nil {
		return nil, err
	}

	if err := fh.kafkaProducer.Produce(msg); err != nil {
		return nil, err
	}

	return &UploadFileResponse{Name: name}, nil
}

func (fh *FileHandler) DownloadFile(ctx context.Context, req *DownloadFileRequest) (*DownloadFileResponse, error) {
	obj, err := fh.svc.DownloadFile(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &DownloadFileResponse{Name: req.Name, Data: obj}, nil
}
