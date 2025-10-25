package grpcclient

import (
	"google.golang.org/grpc"

	pbf "github.com/cdxy1/go-file-storage/internal/grpc/file"
	pbm "github.com/cdxy1/go-file-storage/internal/grpc/metadata"
	"google.golang.org/grpc/credentials/insecure"
)

func NewFileGrpcClient() (pbf.FileServiceClient, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pbf.NewFileServiceClient(conn)
	return client, nil
}

func NewMetadataGprcClient() (pbm.MetadataServiceClient, error) {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pbm.NewMetadataServiceClient(conn)
	return client, err
}
