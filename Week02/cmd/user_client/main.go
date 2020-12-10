package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/stanleyhsu/Go-000/Week02/proto/user/v1"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "stanley"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUserByName(ctx, &pb.UserRequest{Name: name})
	if err != nil {
		log.Fatalf("could not get user info: %+v", err)
	}
	log.Printf("User info: %+v", r)
}
