package metadata

import (
	"net"

	"github.com/cdxy1/go-file-storage/internal/grpc/metadata"
	"github.com/cdxy1/go-file-storage/internal/repo"
	"github.com/cdxy1/go-file-storage/internal/service"
	"google.golang.org/grpc"
)

func NewApp() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic("metadata grpc server not started")
	}

	r, err := repo.NewMetadataRepo()
	if err != nil {
		panic("postgres error")
	}

	svc := service.NewMetadataService(r)

	handler := metadata.NewMetadataHandler(svc)

	grpcSrv := grpc.NewServer()

	metadata.RegisterMetadataServiceServer(grpcSrv, handler)

	if err := grpcSrv.Serve(lis); err != nil {
		panic("grpc server not started")
	}
}
