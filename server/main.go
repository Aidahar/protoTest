package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "protTest/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

type server struct {
	pb.UnimplementedMessagerServer
}

func (s *server) GetMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Println("Request")
	return &pb.MessageResponse{SessionId: "HELLO MOTHERFUCKER!", Frequency: 66.66, Data: timestamppb.Now()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterMessagerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
