package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	pb "study/stream/bidirectionalStream/echo"

	"google.golang.org/grpc"
)

const port = "localhost:1234"

type Echo struct {
	pb.UnimplementedEchoServer
}

// BidirectionalStreamingEcho 双向流服务端
/*
// 1. 建立连接 获取client
// 2. 通过client调用方法获取stream
// 3. 开两个goroutine（使用 chan 传递数据） 分别用于Recv()和Send()
// 3.1 一直Recv()到err==io.EOF(即客户端关闭stream)
// 3.2 Send()则自己控制什么时候Close 服务端stream没有close 只要跳出循环就算close了。 具体见https://github.com/grpc/grpc-go/issues/444
*/
func (echo *Echo) BidirectionalStreamingEcho(stream pb.Echo_BidirectionalStreamingEchoServer) error {
	var (
		waitGroup sync.WaitGroup
		msgCh = make(chan string)
	)
	waitGroup.Add(1)
	go func ()  {
		defer waitGroup.Done()
		for v := range msgCh {
			err := stream.Send(&pb.EchoResponse{Message: v})
			if err != nil {
				fmt.Println("send err", err)
			}
			continue
		}
	}()

	waitGroup.Add(1)
	go func ()  {
		defer waitGroup.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("recv error:%v", err)
			}
			fmt.Printf("Recved :%v \n", req.GetMessage())
			msgCh <- req.GetMessage()
		}
		close(msgCh)
		fmt.Println("close chanel")
	}()
	waitGroup.Wait()
	return nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &Echo{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("listen err: %v", err)
	}
}