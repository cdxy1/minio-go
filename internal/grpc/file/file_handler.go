package file

import (
	"context"

	"github.com/cdxy1/go-file-storage/internal/repo"
)

type FileHandler struct {
	UnimplementedFileServiceServer
	svc *repo.File
}

func NewFileHandler(svc *repo.File) *FileHandler {
	return &FileHandler{svc: svc}
}

func (fh *FileHandler) UploadFile(ctx context.Context, req *UploadFileRequest) (*UploadFileResponse, error) {
	// fh.svc.Put(ctx, req.Name, io.)

	return &UploadFileResponse{Name: "fgsddfs"}, nil
}
