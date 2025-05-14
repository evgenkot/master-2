package main

import (
	"context"
	"log"
	"net"

	"reverse-service/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedReverseServiceServer
}

func (s *server) Reverse(ctx context.Context, in *proto.ReverseRequest) (*proto.ReverseResponse, error) {
	return &proto.ReverseResponse{Output: reverseString(in.Input)}, nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterReverseServiceServer(s, &server{})
	log.Printf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
