package main

import (
	"io"
	"log"
	"context"
	pb "study/stream/serverStream/echo"

	"google.golang.org/grpc"
)

const address = "localhost:50051"

/*
1. 建立连接 获取client
2. 通过 client 获取stream
3. for循环中通过stream.Recv()依次获取服务端推送的消息
4. err==io.EOF则表示服务端关闭stream了
*/
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewEchoClient(conn)
	// 2.调用获取stream
	stream, err := client.ServerStreamingEcho(context.Background(), &pb.EchoRequest{Message: "Hello World"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}

	// 3. for循环获取服务端推送的消息
	for {
		// 通过 Recv() 不断获取服务端send()推送的消息
		resp, err := stream.Recv()
		// 4. err==io.EOF则表示服务端关闭stream了 退出
		if err == io.EOF {
			log.Println("server closed")
			break
		}
		if err != nil {
			log.Printf("Recv error:%v", err)
			continue
		}
		log.Printf("Recv data:%v", resp.GetMessage())
	}
}