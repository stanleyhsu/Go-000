package main

import (
	"context"
	"log"
	"net"

	"github.com/stanleyhsu/Go-000/Week02/internal/service"

	pb "github.com/stanleyhsu/Go-000/Week02/proto/user"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement user.GetUserByName.
type server struct {
	pb.UnimplementedUserServer
}

func (*server) GetUserByName(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	log.Printf("Received: %v", in.GetName())
	rsp, err := service.GetUserByName(ctx, in.GetName())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &pb.UserReply{Name: rsp.Name, Age: rsp.Age, Cellphone: rsp.Cellphone, Email: rsp.Email}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
