package config

import (
	"fmt"
	"sync/atomic"

	"github.com/BurntSushi/toml"
)

const configPath = "configs/grpc.toml"

type GrpcServerConfig struct {
	Server ServerConf `toml:"Server"`
}
type ServerConf struct {
	Addr string `toml:"addr"`
}

var _config atomic.Value

func loadConfigFromFile(filepath string) GrpcServerConfig {
	var grpcConf GrpcServerConfig
	if _, err := toml.DecodeFile(filepath, &grpcConf); err != nil {
		fmt.Printf("Parse config file:%s error:%s", filepath, err)
		panic(err)
	}
	return grpcConf
}
func Init() {
	_config.Store(loadConfigFromFile(configPath))
}

func GetGrpcServerConfig() ServerConf {
	conf := _config.Load().(GrpcServerConfig)
	return conf.Server
}
