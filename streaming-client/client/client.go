package main

import (
	"context"
	"log"
	"time"

	pb "github.com/wagaru/microservice/streaming-client/gen/pb-go/helloworld"

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
	stream, err := c.SayHello(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for i := 0; i <= 10; i++ {
		stream.Send(&pb.HelloRequest{Name: defaultName})
		time.Sleep(100 * time.Millisecond)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("receive faild: %v", err)
	}
	log.Printf("Greeting: %s", resp.GetMessage())
}
