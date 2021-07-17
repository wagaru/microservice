package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/wagaru/microservice/streaming-server/gen/pb-go/helloworld"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50052"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	stream, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream read failed: %v", err)
		}
		log.Println(reply.GetMessage())
	}
}
