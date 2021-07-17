package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/wagaru/microservice/streaming-bidirectional/gen/pb-go/helloworld"

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

	stream, err := c.SayHello(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			reply, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("failed to receive:%v", err)
			}
			log.Println(reply.GetMessage())
		}
	}()
	for i := 0; i <= 10; i++ {
		if err := stream.Send(&pb.HelloRequest{Name: defaultName}); err != nil {
			log.Fatalf("error while sending:%v", err)
		}
		time.Sleep(time.Second)
	}
	log.Printf("Close Send")
	stream.CloseSend()
	<-waitc
}
