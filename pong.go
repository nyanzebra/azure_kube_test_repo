package main

import (
	"context"
	"net"

	"log"

	"google.golang.org/grpc"

	message "aktr/message"
)

// RealMessageServer implements MessageServer
type RealMessageServer struct {
	message.UnimplementedMessengerServer
}

// Ping responder
func (*RealMessageServer) Ping(ctx context.Context, request *message.PingRequest) (*message.PongResponse, error) {
	log.Printf("Message: %v\n", request.Message)
	return &message.PongResponse{Message: "Sup"}, nil
}

func (*RealMessageServer) mustEmbedUnimplementedMessengerServer() {}

func main() {
	lis, err := net.Listen("tcp", "localhost:9876")
	if err != nil {
		log.Fatalf("Err: %v\n", err)
	}
	server := grpc.NewServer()

	message.RegisterMessengerServer(server, &RealMessageServer{})
	server.Serve(lis)
}
