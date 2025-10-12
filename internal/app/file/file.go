package file

import (
	"net"

	"github.com/cdxy1/go-file-storage/internal/grpc/file"
	"github.com/cdxy1/go-file-storage/internal/repo"
	"google.golang.org/grpc"
)

func NewApp() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("Grpc server not started")
	}

	svc, err := repo.NewFileRepo()

	if err != nil {
		panic("Grpc server not started")
	}

	handler := file.NewFileHandler(svc)

	grpcSrv := grpc.NewServer()

	file.RegisterFileServiceServer(grpcSrv, handler)

	if err := grpcSrv.Serve(lis); err != nil {
		panic("grpc server not started")
	}
}
