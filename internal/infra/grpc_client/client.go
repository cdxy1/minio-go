package grpcclient

import (
	"google.golang.org/grpc"

	pb "github.com/cdxy1/go-file-storage/internal/grpc/file"
)

func NewFileGrpcClient() (pb.FileServiceClient, error){
	conn, err := grpc.NewClient("localhost:50051")

	if err != nil {
		return nil, err
	}

	client := pb.NewFileServiceClient(conn)
	return client, nil
}
