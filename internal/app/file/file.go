package file

import (
	"net"

	"google.golang.org/grpc"

	"github.com/cdxy1/go-file-storage/internal/grpc/file"
	"github.com/cdxy1/go-file-storage/internal/repo"
	"github.com/cdxy1/go-file-storage/internal/service"
)

func NewApp() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("grpc server not started")
	}

	r, err := repo.NewFileRepo()
	if err != nil {
		panic("Grpc server not started")
	}

	svc := service.NewFileService(r)

	handler := file.NewFileHandler(svc)

	grpcSrv := grpc.NewServer()

	file.RegisterFileServiceServer(grpcSrv, handler)

	if err := grpcSrv.Serve(lis); err != nil {
		panic("grpc server not started")
	}
}
