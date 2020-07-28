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
	"sync"
)

var (
	port        = flag.Int("port", 10000, "The server port")
	server_type = flag.String("chat_type", "cs", "Type of chatting 'group or cs'")
)

type chatServer struct {
	mu      sync.Mutex // protects streams
	streams []*Chat_GroupChatServer
	UnimplementedChatServer
}

func (s *chatServer) ClientServerChat(stream Chat_ClientServerChatServer) error {
	fmt.Printf("Client:Start Chat\n")
	go func() error {
		defer func() {
			fmt.Printf("Client:End Chat\n")
		}()
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
	// 多个client时会有问题，相当于多个gorutine同时阻塞在此，消息随机发送给客户端
	// 实现群发需要共享同一个终端
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

func (s *chatServer) GroupChat(stream Chat_GroupChatServer) error {
	fmt.Printf("Client:Start Chat\n")
	s.mu.Lock()
	s.streams = append(s.streams, &stream)
	fmt.Printf("Current connection:%d\n", len(s.streams))
	s.mu.Unlock()
	defer func() {
		fmt.Printf("Client:End Chat\n")
		index := -1
		for i := 0; i < len(s.streams); i++ {
			if s.streams[i] == &stream {
				index = i
			}
		}
		if index >= 0 {
			s.streams = append(s.streams[:index], s.streams[index+1:]...)
		}
		fmt.Printf("Current connection:%d\n", len(s.streams))
	}()
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
		s.mu.Lock()
		for _, s := range s.streams {
			if s == &stream {
				continue
			}
			if err := (*s).Send(in); err != nil {
				return err
			}
		}
		s.mu.Unlock()
	}
}

//自己定义的
func newServer(server_type string) *chatServer {
	s := &chatServer{}
	if server_type == "group" {
		go func() error {
			scanner := bufio.NewScanner(os.Stdin)
			for {
				if scanner.Scan() {
					line := scanner.Text()
					info := &ChatInfo{Content: line}
					fmt.Printf("Server:%s\n", line)
					s.mu.Lock()
					for _, stream := range s.streams {
						if err := (*stream).Send(info); err != nil {
							return err
						}
					}
					s.mu.Unlock()
				}
			}
		}()
	}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	grpcServer := grpc.NewServer()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	RegisterChatServer(grpcServer, newServer(*server_type))
	grpcServer.Serve(lis)
}
