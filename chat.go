package chat

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) BroadcastMessage(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %s", in.Body)
	return &Message{Body: "Broadcasted message!"}, nil
}

func (s *Server) CalculateFactorial(ctx context.Context, req *NumberRequest) (*CalculationResponse, error) {
	n := req.Number
	if n < 0 {
		return nil, fmt.Errorf("negative numbers not allowed")
	}

	result := int64(1)
	for i := int32(2); i <= n; i++ {
		result *= int64(i)
	}
	return &CalculationResponse{Result: result}, nil
}
