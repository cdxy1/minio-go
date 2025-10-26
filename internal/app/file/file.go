package file

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/cdxy1/go-file-storage/internal/config"
	"github.com/cdxy1/go-file-storage/internal/grpc/file"
	"github.com/cdxy1/go-file-storage/internal/infra/kafka/producer"
	"github.com/cdxy1/go-file-storage/internal/repo"
	"github.com/cdxy1/go-file-storage/internal/service"
	"github.com/cdxy1/go-file-storage/pkg/logger"
)

func NewApp() {
	cfg := config.GetConfig()
	log := logger.SetupLogger()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.File.Port))
	if err != nil {
		log.Error("can not able to run file grpc server", "error", err)
		os.Exit(1)
	}

	r, err := repo.NewFileRepo(log)
	if err != nil {
		log.Error("error in file repo", "error", err)
		os.Exit(1)
	}

	svc := service.NewFileService(r, log)
	kafka, err := producer.NewProducer()
	if err != nil {
		log.Error("unable to run kafka producer", "error", err)
		os.Exit(1)
	}

	handler := file.NewFileHandler(svc, kafka, log)

	grpcSrv := grpc.NewServer()

	file.RegisterFileServiceServer(grpcSrv, handler)

	if err := grpcSrv.Serve(lis); err != nil {
		log.Error("unable to serve file grpc server", "error", err)
		os.Exit(1)
	}
}
