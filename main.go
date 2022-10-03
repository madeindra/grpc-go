package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/madeindra/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedTodoServiceServer
}

func (s *server) Create(ctx context.Context, in *pb.TodoRequest) (*pb.TodoResponse, error) {
	fmt.Println("Got ID ", in.Content)
	fmt.Println("Got Title ", in.Title)

	return &pb.TodoResponse{}, nil
}

func (s *server) Read(ctx context.Context, in *emptypb.Empty) (*pb.TodoList, error) {
	return &pb.TodoList{}, nil
}

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterTodoServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}
