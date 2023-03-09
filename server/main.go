package main

import (
	"context"
	"log"
	"net"
	pb "store/proto"

	"store/storage/postgres"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StoreServer struct {
	pb.UnimplementedStoreServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listening: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterStoreServiceServer(s, &StoreServer{})

	log.Printf("server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *StoreServer) CreateStore(ctx context.Context, in *pb.Store) (*pb.Store, error) {
	log.Printf("received: %v", in.GetName())
	store, err := postgres.CreateStore(&pb.Store{
		Id:          1,
		Name:        in.Name,
		Description: in.Description,
		Addresses:   in.Addresses,
		IsOpen:      in.IsOpen,
	})
	if err != nil {
		log.Fatalf("failed to create store in server: %v", err)
	}

	return store, nil
}

func (s *StoreServer) GetStore(ctx context.Context, in *pb.GetStoreRequest) (*pb.Store, error) {
	store, err := postgres.GetStore(in.Id)
	if err != nil {
		log.Fatalf("failed to get store in server: %v", err)
	}

	return store, nil
}

func (s *StoreServer) UpdateStore(ctx context.Context, in *pb.Store) (*emptypb.Empty, error) {
	log.Printf("updated: %v", in.GetName())
	err := postgres.UpdateStore(in)
	if err != nil {
		log.Fatalf("failed to update store in server: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *StoreServer) DeleteStore(ctx context.Context, in *pb.GetStoreRequest) (*emptypb.Empty, error) {
	err := postgres.DeleteStore(in.Id)
	if err != nil {
		log.Fatalf("failed to delete store in server: %v", err)
	}

	return &emptypb.Empty{}, nil
}
