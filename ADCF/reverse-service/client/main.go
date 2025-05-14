package main

import (
	"context"
	"flag"
	"log"
	"time"

	"reverse-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := flag.String("server", "localhost:50051", "Server address (host:port)")
	input := flag.String("input", "", "Input string to reverse")
	flag.Parse()

	if *input == "" {
		log.Fatal("Input string is required")
	}

	// Connect to server
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create client and make request
	client := proto.NewReverseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Reverse(ctx, &proto.ReverseRequest{Input: *input})
	if err != nil {
		log.Fatalf("Reverse failed: %v", err)
	}

	log.Printf("Reversed: %s", resp.Output)
}
