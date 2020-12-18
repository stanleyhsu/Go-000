package grpc

import (
	"fmt"
	"net"

	"github.com/stanleyhsu/Go-000/Week04/internal/config"

	pb "github.com/stanleyhsu/Go-000/Week04/api/hello/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	lis net.Listener
	g   *grpc.Server
}

func New(svc pb.GreeterServer, conf config.ServerConf) *Server {
	listener, err := net.Listen("tcp", fmt.Sprintf(conf.Addr))
	if err != nil {
		fmt.Printf("net listen tcp addr:%s error:%s", conf.Addr, err.Error())
		return nil
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, svc)
	reflection.Register(grpcServer)
	server := &Server{
		lis: listener,
		g:   grpcServer,
	}
	return server
}

func (s *Server) Serve() error {
	fmt.Println("Try to start grpc on addr", s.lis.Addr())
	if err := s.g.Serve(s.lis); err != nil {
		return err
	}
	return nil
}
