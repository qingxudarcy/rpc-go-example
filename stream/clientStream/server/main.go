package main

import (
	"io"
	"log"
	"net"
	pb "study/stream/clientStream/echo"

	"google.golang.org/grpc"
)

const port = "localhost:50051"

type Echo struct {
	pb.UnimplementedEchoServer
}

func (echo *Echo) ClientStreamingEcho(stream pb.Echo_ClientStreamingEchoServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("client closed")
			return stream.SendAndClose(&pb.EchoResponse{Message: "ok"})
		}
		if err != nil {
			return err
		}
		log.Printf("Recved %s", req.GetMessage())
	}
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &Echo{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}