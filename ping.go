package main

import (
	"context"
	"log"

	"time"

	"google.golang.org/grpc"

	message "aktr/message"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:9876", opts...)
	if err != nil {
		log.Fatalf("Err: %v\n", err)
	}
	defer conn.Close()

	client := message.NewMessengerClient(conn)

	for {
		req := &message.PingRequest{Message: "Hi"}
		pong, err := client.Ping(context.Background(), req)
		if err != nil {
			log.Fatalf("Err: %v\n", err)
			continue
		}
		log.Printf("Message: %v\n", pong.Message)
		time.Sleep(2 * time.Second)
	}
}
