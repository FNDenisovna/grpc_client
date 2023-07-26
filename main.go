package main

import (
	"context"
	"fmt"
	"grpcclient/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {

	//var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Printf("Can't connect to the grpc-server! %v\n", err)
	}
	defer conn.Close()

	testClient := NewAlbumServiceClient(conn)
	if err = testClient.Run(context.Background()); err != nil {
		log.Printf("Grpc-server has internal error during a testcase! %v\n", err)
	} else {
		log.Printf("Testcase is finished successful!\n")
	}
}

type AlbumServiceClient struct {
	service proto.GrpcAlbumClient
}

func NewAlbumServiceClient(conn grpc.ClientConnInterface) *AlbumServiceClient {
	return &AlbumServiceClient{
		service: proto.NewGrpcAlbumClient(conn),
	}
}

func (s *AlbumServiceClient) Run(ctx context.Context) error {
	return s.GetAlbums(ctx)
}

func (s *AlbumServiceClient) GetAlbums(ctx context.Context) error {
	resp, err := s.service.GetAlbums(ctx, &proto.GetAlbumsRequest{
		Limit: 10,
	})
	if err != nil {
		return fmt.Errorf("Error in method GetAlbums: %v\n", err)
	} else {
		log.Printf("Response from grpc-server: %s\n", resp)
		return nil
	}
}
