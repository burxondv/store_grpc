package main

import (
	"context"
	"log"
	"time"

	pb "store/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Store struct {
	Id          int64
	Name        string
	Description string
	Addresses   []string
	IsOpen      bool
}

const (
	ServerAddress = "localhost:8000"
)

func main() {
	conn, err := grpc.Dial(ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	c := pb.NewStoreServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// CREATE STORE
	/*
		store, err := c.CreateStore(ctx, &pb.Store{
			Name:        "Thomas",
			Description: "32 years old",
			Addresses: []string{
				"California",
				"Los Angeles",
			},
			IsOpen: true,
		})
		if err != nil {
			log.Fatalf("failed to create store: %v", err)
		}
		fmt.Println(store)
	*/

	// GET STORE
	/*
		store, err := c.GetStore(ctx, &pb.GetStoreRequest{
			Id: 2,
		})
		if err != nil {
			log.Fatalf("failed to get store in client: %v", err)
		}
		fmt.Println(store)
	*/

	// UPDATE STORE
	/*
	_, err = c.UpdateStore(ctx, &pb.Store{})
	if err != nil {
		log.Fatalf("failed to update store in client: %v", err)
	}
	*/

	// DELETE STORE
	_, err = c.DeleteStore(ctx, &pb.GetStoreRequest{
		Id: 8,
	})
	if err != nil {
		log.Fatalf("failed to delete store in client: %v", err)
	}

}
