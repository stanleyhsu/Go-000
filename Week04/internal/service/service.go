package service

import (
	"context"

	"github.com/google/wire"

	"github.com/stanleyhsu/Go-000/Week04/internal/biz"
	"github.com/stanleyhsu/Go-000/Week04/internal/data"

	pb "github.com/stanleyhsu/Go-000/Week04/api/hello/v1"
)

var ProviderSet = wire.NewSet(
	data.Provider,
	wire.Bind(new(biz.Messager), new(*data.DataMessage)),
	New,
)

// Service service.
type Service struct {
	messager biz.Messager
	pb.UnimplementedGreeterServer
}

func New(m biz.Messager) *Service {
	return &Service{
		messager: m,
	}
}

func (s *Service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	msg := s.messager.Get()
	return &pb.HelloReply{Message: "Hello " + in.GetName() + msg}, nil
}
