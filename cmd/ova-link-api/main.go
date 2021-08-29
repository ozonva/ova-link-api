package main

import (
	"log"
	"net"

	"github.com/ozonva/ova-link-api/internal/api"

	linkAPI "github.com/ozonva/ova-link-api/pkg/ova-link-api"

	"google.golang.org/grpc"
)

func main() {
	const grpcPort = ":82"
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	linkAPI.RegisterLinkAPIServer(s, api.NewLinkAPI())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return
}
