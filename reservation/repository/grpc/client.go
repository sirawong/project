package grpc

import (
	"log"

	"google.golang.org/grpc"
)

func (grpcRepo *Config) NewClient() *grpc.ClientConn {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(grpcRepo.Port, opts...)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}
