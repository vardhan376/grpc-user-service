package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "protobuf_project/proto"

	"google.golang.org/grpc"
)

// Server struct
type server struct {
	pb.UnimplementedUserServiceServer
}

// Server-side Go struct (DB/service representation)
type UserStruct struct {
	Uid   string
	Name  string
	Age   int32
	Email string
	Phone string
	City  string
}

// Implement GetUserData
func (s *server) GetUserData(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Printf("✅ Client requested UID: %s, fields: %v\n", req.Uid, req.Fields)

	// Full user struct (as if fetched from DB)
	userStruct := UserStruct{
		Uid:   req.Uid,
		Name:  "Harsh Vardhan",
		Age:   23,
		Email: "harsh@example.com",
		Phone: "9876543210",
		City:  "Hyderabad",
	}

	// Populate only requested fields into proto User
	userProto := &pb.User{Uid: userStruct.Uid} // UID hamesha include

	for _, field := range req.Fields {
		switch field {
		case "name":
			userProto.Name = userStruct.Name
		case "age":
			userProto.Age = userStruct.Age
		case "email":
			userProto.Email = userStruct.Email
		case "phone":
			userProto.Phone = userStruct.Phone
		case "city":
			userProto.City = userStruct.City
		}
	}

	return &pb.UserResponse{User: userProto}, nil
}

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("🚀 Server started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
