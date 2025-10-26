package file

import (
	"context"
	"log/slog"

	"github.com/cdxy1/go-file-storage/internal/infra/kafka/producer"
	"github.com/cdxy1/go-file-storage/internal/lib"
	"github.com/cdxy1/go-file-storage/internal/service"
)

type FileHandler struct {
	UnimplementedFileServiceServer
	svc           *service.FileService
	kafkaProducer *producer.Producer
	logger        *slog.Logger
}

func NewFileHandler(svc *service.FileService, kafka *producer.Producer, logger *slog.Logger) *FileHandler {
	return &FileHandler{svc: svc, kafkaProducer: kafka, logger: logger}
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
		fh.logger.Error("error while producing message", "message", req.Name)
		return nil, err
	}
	fh.logger.Info("message successfully produced", "message", req.Name)

	return &UploadFileResponse{Name: name}, nil
}

func (fh *FileHandler) DownloadFile(ctx context.Context, req *DownloadFileRequest) (*DownloadFileResponse, error) {
	obj, err := fh.svc.DownloadFile(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &DownloadFileResponse{Name: req.Name, Data: obj}, nil
}
