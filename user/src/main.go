package main

import (
	"context"
	"log"
	"net"

	pb "user/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users map[string]string // Simple in-memory user storage: username -> password
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	if _, exists := s.users[req.Username]; exists {
		return nil, grpc.Errorf(409, "User already exists")
	}
	s.users[req.Username] = req.Password
	return &pb.UserResponse{UserId: req.Username, Token: "dummy-token"}, nil
}

func (s *server) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.UserResponse, error) {
	if pwd, exists := s.users[req.Username]; exists && pwd == req.Password {
		return &pb.UserResponse{UserId: req.Username, Token: "dummy-token"}, nil
	}
	return nil, grpc.Errorf(401, "Invalid credentials")
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{users: make(map[string]string)})

	log.Println("User Service is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
