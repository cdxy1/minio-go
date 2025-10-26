package metadata

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/cdxy1/go-file-storage/internal/config"
	"github.com/cdxy1/go-file-storage/internal/grpc/metadata"
	"github.com/cdxy1/go-file-storage/internal/infra/kafka/consumer"
	"github.com/cdxy1/go-file-storage/internal/repo"
	"github.com/cdxy1/go-file-storage/internal/service"
)

func NewApp() {
	cfg := config.GetConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Metadata.Port))
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

	cons, err := consumer.NewConsumer(handler)
	if err != nil {
		panic("kafka not working")
	}
	go cons.Start()

	if err := grpcSrv.Serve(lis); err != nil {
		panic("grpc server not started")
	}
}
