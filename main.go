// to run the server use the below command
// 1. go run ./server/main.go
// 2. protoc --go_out=. --go-grpc_out=. chat.proto
// 3. go run ./client/client.go

package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_testing/chat"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		fmt.Println("success") // success message
	}

	grpcServer := grpc.NewServer()
	s := chat.Server{}

	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
