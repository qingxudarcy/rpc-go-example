package main

import (
	"log"
	"context"

	pb "study/stream/clientStream/echo"

	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewEchoClient(conn)
	stream, err := client.ClientStreamingEcho(context.Background())
	if err != nil {
		log.Fatalf("Sum() err: %v", err)
	}
	for i := int64(0); i < 2; i++ {
		err := stream.Send(&pb.EchoRequest{Message: "hello world"})
		if err != nil {
			log.Printf("send error: %v", err)
			continue
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("CloseAndRecv() error: %v", err)
	}
	log.Printf("sum: %s", resp.GetMessage())
}