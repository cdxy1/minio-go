package metadata

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/cdxy1/go-file-storage/internal/config"
	"github.com/cdxy1/go-file-storage/internal/grpc/metadata"
	"github.com/cdxy1/go-file-storage/internal/infra/kafka/consumer"
	"github.com/cdxy1/go-file-storage/internal/repo"
	"github.com/cdxy1/go-file-storage/internal/service"
	"github.com/cdxy1/go-file-storage/pkg/logger"
)

func NewApp() {
	cfg := config.GetConfig()
	log := logger.SetupLogger()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Metadata.Port))
	if err != nil {
		log.Error("can not able to run metadata grpc server", "error", err)
		os.Exit(1)
	}

	r, err := repo.NewMetadataRepo(log)
	if err != nil {
		log.Error("can not able to connect postgres", "error", err)
		os.Exit(1)
	}

	svc := service.NewMetadataService(r, log)

	handler := metadata.NewMetadataHandler(svc)

	grpcSrv := grpc.NewServer()

	metadata.RegisterMetadataServiceServer(grpcSrv, handler)

	cons, err := consumer.NewConsumer(handler, log)
	if err != nil {
		log.Error("unable to run kafka consumer")
		os.Exit(1)
	}
	go cons.Start()

	if err := grpcSrv.Serve(lis); err != nil {
		log.Error("unable to serve metadata grpc server", "error", err)
		os.Exit(1)
	}
}
