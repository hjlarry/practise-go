package main

import (
	"io"
	"log"
	"net"

	pb "../../proto"
	"google.golang.org/grpc"
)

type StreamService struct{}

const PORT = "50051"

func main() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	_ = server.Serve(lis)
}

//服务端流式，客户端发起一次RPC请求，服务端通过流式相应多次发送数据集
func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	for n := 0; n < 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

//客户端流式，客户端通过多次RPC请求给服务端，服务端发起一次响应给客户端
func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{
				Pt: &pb.StreamPoint{
					Name:  "gRpc Stream Server:Record",
					Value: 1,
				},
			})
		}
		if err != nil {
			return err
		}
		log.Printf("stream Recv:pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
}

//由客户端以流式的方式发起请求，服务端同样以流式的方式响应请求
//首个请求一定是 Client 发起，但具体交互方式（谁先谁后、一次发多少、响应多少、什么时候关闭）根据程序编写的方式来确定（可以结合协程）
func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "gRPC Client Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++
		log.Printf("stream Recv:pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
}
