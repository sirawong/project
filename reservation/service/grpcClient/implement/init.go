package implement

import (
	"reservation/repository/grpc"
	"reservation/service/grpcClient"
	auth "reservation/service/grpcClient/protobuf/auth"
	cinema "reservation/service/grpcClient/protobuf/cinema"
)

type implementation struct {
	pbAuth   auth.AuthenticationClient
	pbCinema cinema.CinemaServiceClient
}

func New(grpcAuthRepo grpc.RepositoryGRPC, grpcCinemaRepo grpc.RepositoryGRPC) (service grpcClient.Service) {
	connAuth := grpcAuthRepo.NewClient()
	connCinema := grpcCinemaRepo.NewClient()

	impl := &implementation{
		pbAuth:   auth.NewAuthenticationClient(connAuth),
		pbCinema: cinema.NewCinemaServiceClient(connCinema),
	}

	return impl
}
