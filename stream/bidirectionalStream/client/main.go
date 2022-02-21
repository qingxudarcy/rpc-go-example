package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	pb "study/stream/bidirectionalStream/echo"

	"google.golang.org/grpc"
)

const address = "localhost:1234"


func bidirectionalStream(stream pb.Echo_BidirectionalStreamingEchoClient) {
	var waitGroup sync.WaitGroup
	
	waitGroup.Add(1)
	go func ()  {
		defer waitGroup.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Sprintf("Recive err: %v", err)
			}
			fmt.Printf("Recive %s \n", req.GetMessage())
		}
	}()

	waitGroup.Add(1)
	go func ()  {
		defer waitGroup.Done()
		for i :=0; i < 2; i ++ {
			err := stream.Send(&pb.EchoRequest{Message: "hello world"})
			if err != nil {
				fmt.Printf("send err: %v", err)
			}
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		fmt.Println("关闭发送")
		if err != nil {
			log.Printf("Send error:%v\n", err)
			return
		}
	}()
	waitGroup.Wait()
	fmt.Println("success")
}



func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connect err: %v", err)
	}
	defer conn.Close()
	client := pb.NewEchoClient(conn)
	stream, err := client.BidirectionalStreamingEcho(context.Background())
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	bidirectionalStream(stream)
}