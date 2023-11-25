package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"playground/agify"
)

type AgifyServer struct {
	agify.UnimplementedAgifyServer
}

func (s *AgifyServer) GetEstimatedAge(ctx context.Context, person *agify.Person) (*agify.Age, error) {
	age := &agify.Age{Age: 15}
	return age, nil
}

func (s *AgifyServer) GetCount(ctx context.Context, person *agify.Person) (*agify.Count, error) {
	count := &agify.Count{Count: 114}
	return count, nil
}

func main() {
	fmt.Printf("start server\n")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	agify.RegisterAgifyServer(s, &AgifyServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Printf("close server\n")
}
