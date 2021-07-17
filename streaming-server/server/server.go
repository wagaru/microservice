package main

import (
	"log"
	"net"

	pb "github.com/wagaru/microservice/streaming-server/gen/pb-go/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(req *pb.HelloRequest, srv pb.Greeter_SayHelloServer) error {
	log.Printf("Received: %v", req.GetName())
	for i := 0; i < 10; i++ {
		srv.Send(&pb.HelloReply{Message: req.GetName()})
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", port)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
