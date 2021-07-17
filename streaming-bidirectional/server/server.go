package main

import (
	"io"
	"log"
	"net"

	pb "github.com/wagaru/microservice/streaming-bidirectional/gen/pb-go/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(srv pb.Greeter_SayHelloServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			log.Print("receive end")
			return nil
		}
		if err != nil {
			log.Printf("receive error:%v", err)
		}
		log.Printf("receive: %v", req.GetName())
		srv.Send(&pb.HelloReply{Message: "hello " + req.GetName()})
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
