package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/wagaru/microservice/streaming-client/gen/pb-go/helloworld"
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
		msg, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&pb.HelloReply{Message: "hello"})
		}
		if err != nil {
			log.Fatalf("could not receive srv: %v", err)
		}
		fmt.Println(msg.GetName())
	}
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
