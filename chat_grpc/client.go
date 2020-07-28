package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

var (
	serverAddr  = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	client_type = flag.String("chat_type", "cs", "Type of chatting 'group or cs'")
)

func runCSChatClient(client ChatClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 900*time.Second)
	defer cancel()
	stream, err := client.ClientServerChat(ctx)
	defer stream.CloseSend()
	if err != nil {
		log.Fatalf("%v.ClientServerChat(_) = _, %v", client, err)
	}
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a info : %v", err)
			}
			log.Printf("Server:%s", in.Content)
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			info := &ChatInfo{Content: line}
			fmt.Printf("Client:%s\n", line)
			if err := stream.Send(info); err != nil {
				log.Fatalf("Failed to send a info: %v", err)
			}
		}
	}
}

func runGroupChatClient(client ChatClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 900*time.Second)
	defer cancel()
	stream, err := client.GroupChat(ctx)
	defer stream.CloseSend()
	if err != nil {
		log.Fatalf("%v.GroupChat(_) = _, %v", client, err)
	}
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a info : %v", err)
			}
			log.Printf("Server:%s", in.Content)
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			info := &ChatInfo{Content: line}
			fmt.Printf("Client:%s\n", line)
			if err := stream.Send(info); err != nil {
				log.Fatalf("Failed to send a info: %v", err)
			}
		}
	}
}

func main() {

	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := NewChatClient(conn)
	if *client_type == "cs" {
		runCSChatClient(client)
	} else {
		runGroupChatClient(client)
	}

}
