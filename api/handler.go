package api

import (
	"golang.org/x/net/context"
	"log"
)

// Server represent the gRPC server
type Server struct {

}

// SayHello generates response to a ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage)(*PingMessage, error){
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}