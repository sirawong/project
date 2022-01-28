package grpc

import (
	"net"

	"google.golang.org/grpc"

	"cinema/config"
	cinemaService "cinema/service/cinema"
	"cinema/service/grpc/protobuf/cinema"
)

type CinemaServer struct {
	cinemaSrv cinemaService.CinemaService
}

func NewServer(cinemaSrv cinemaService.CinemaService, appConfig *config.Config) {

	server := CinemaServer{
		cinemaSrv: cinemaSrv,
	}

	lis, _ := net.Listen("tcp", appConfig.GRPCHost)

	grpcServer := grpc.NewServer()

	cinema.RegisterCinemaServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
