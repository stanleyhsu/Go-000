// +build wireinject

package server

import (
	"github.com/google/wire"
	pb "github.com/stanleyhsu/Go-000/Week04/api/hello/v1"
	"github.com/stanleyhsu/Go-000/Week04/internal/config"
	grpcServer "github.com/stanleyhsu/Go-000/Week04/internal/server/grpc"
	"github.com/stanleyhsu/Go-000/Week04/internal/service"
)

func InitApp() *App {
	panic(wire.Build(service.ProviderSet, wire.Bind(new(pb.GreeterServer), new(*service.Service)),
		config.GetGrpcServerConfig,
		grpcServer.New, NewApp))
}
