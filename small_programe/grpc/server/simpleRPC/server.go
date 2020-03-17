package main

import (
	"context"
	"log"
	"net"

	pb "../../proto"
	"google.golang.org/grpc"
)

type SearchService struct{}

type HelloService struct {
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + "Server"}, nil
}

func (h *HelloService) Hello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: r.GetName() + "  from Server"}, nil
}

func (h *HelloService) SayHelloAgain(context.Context, *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "abc logic from Server"}, nil
}

const PORT = "50051"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})
	pb.RegisterHelloServiceServer(server, &HelloService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	_ = server.Serve(lis)
}
