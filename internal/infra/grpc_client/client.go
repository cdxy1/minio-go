package grpcclient

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/cdxy1/go-file-storage/internal/config"
	pbf "github.com/cdxy1/go-file-storage/internal/grpc/file"
	pbm "github.com/cdxy1/go-file-storage/internal/grpc/metadata"
	"google.golang.org/grpc/credentials/insecure"
)

func NewFileGrpcClient() (pbf.FileServiceClient, error) {
	cfg := config.GetConfig()

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.File.Host, cfg.File.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pbf.NewFileServiceClient(conn)
	return client, nil
}

func NewMetadataGprcClient() (pbm.MetadataServiceClient, error) {
	cfg := config.GetConfig()

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Metadata.Host, cfg.Metadata.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pbm.NewMetadataServiceClient(conn)
	return client, err
}
