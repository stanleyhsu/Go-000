package main

import (
	"github.com/stanleyhsu/Go-000/Week04/internal/config"
	"github.com/stanleyhsu/Go-000/Week04/internal/server"
)

func main() {
	config.Init()
	s := server.InitApp()
	s.Start()
}
