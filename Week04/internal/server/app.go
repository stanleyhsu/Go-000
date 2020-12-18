package server

import (
	grpcServer "github.com/stanleyhsu/Go-000/Week04/internal/server/grpc"
	"github.com/stanleyhsu/Go-000/Week04/internal/service"
)

type App struct {
	svc  *service.Service
	grpc *grpcServer.Server
}

func NewApp(svc *service.Service, g *grpcServer.Server) (app *App) {
	app = &App{
		svc:  svc,
		grpc: g,
	}
	return app
}

func (app *App) Start() {
	if err := app.grpc.Serve(); err != nil {
		panic(err)
	}
}

//func InitApp() *App {
//	data := data.New()
//	service := service.New(data)
//
//	config := config.GetGrpcServerConfig()
//	grpc := grpcServer.New(service, config)
//
//	app := NewApp(service, grpc)
//	return app
//}
