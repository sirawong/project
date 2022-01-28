package grpc

import "google.golang.org/grpc"

type Config struct {
	Network string
	Port    string
}

type RepositoryGRPC interface {
	NewClient() *grpc.ClientConn
}

func New(config *Config) (grpc *Config) {
	return &Config{config.Network, config.Port}
}
