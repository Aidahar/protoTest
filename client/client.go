package main

import (
	"context"
	"flag"
	"log"
	pb "protTest/gen/go"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:9000", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewMessagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMessage(ctx, &pb.MessageRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Session_id: %s, frequency: %f, date: %v", r.GetSessionId(), r.GetFrequency(), r.GetData().AsTime().Format(time.UnixDate))
}
