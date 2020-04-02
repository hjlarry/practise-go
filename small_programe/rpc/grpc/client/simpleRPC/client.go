package main

import (
	"context"
	"log"

	pb "../../proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &pb.HelloRequest{
		Name: "hehehehei",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetMessage())
}
