package main

import (
	"bufio"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type chatServer struct {
	UnimplementedChatServer
}

func (s *chatServer) BidirectionalStreamingChat(stream Chat_BidirectionalStreamingChatServer) error {
	fmt.Printf("Client:Start Chat\n")
	defer func() {
		fmt.Printf("Client:End Chat\n")
	}()
	go func() error {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			content := in.Content

			log.Printf("Client:%s\n", content)
		}

	}()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			info := &ChatInfo{Content: line}
			fmt.Printf("Server:%s\n", line)
			if err := stream.Send(info); err != nil {
				return err
			}
		}
	}
}

//自己定义的
func newServer() *chatServer {
	s := &chatServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	grpcServer := grpc.NewServer()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	RegisterChatServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
