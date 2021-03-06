// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"github.com/stanleyhsu/Go-000/Week04/internal/config"
	"github.com/stanleyhsu/Go-000/Week04/internal/data"
	"github.com/stanleyhsu/Go-000/Week04/internal/server/grpc"
	"github.com/stanleyhsu/Go-000/Week04/internal/service"
)

// Injectors from wire.go:

func InitApp() *App {
	dataMessage := data.New()
	serviceService := service.New(dataMessage)
	serverConf := config.GetGrpcServerConfig()
	server := grpc.New(serviceService, serverConf)
	app := NewApp(serviceService, server)
	return app
}
