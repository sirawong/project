package grpc

import (
	"net"

	"google.golang.org/grpc"

	"user/config"
	authService "user/service/auth"
	"user/service/grpc/protobuf/auth"
)

type AuthenticationServer struct {
	authSrv authService.Service
}

func NewServer(authSrv authService.Service, appConfig *config.Config) {

	server := AuthenticationServer{
		authSrv: authSrv,
	}

	lis, _ := net.Listen("tcp", appConfig.GRPCHost)

	grpcServer := grpc.NewServer()

	auth.RegisterAuthenticationServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
