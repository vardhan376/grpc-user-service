package main

import (
	"context"
	"log"
	"time"

	pb "protobuf_project/proto"

	"google.golang.org/grpc"
)

func main() {
	// Connect to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("❌ Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Create request
	req := &pb.UserRequest{
		Uid:    "101",
		Fields: []string{"name", "city"},
	}

	// Send request with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.GetUserData(ctx, req)
	if err != nil {
		log.Fatalf("❌ Could not get response: %v", err)
	}

	log.Printf("🎯 User Data: %+v\n", res.User)
}
