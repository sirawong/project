package implement

import (
	"cinema/repository/grpc"
	"cinema/service/grpcClient"
	auth "cinema/service/grpcClient/protobuf/auth"
)

type implementation struct {
	pbAuth auth.AuthenticationClient
}

func New(grpcAuthRepo grpc.RepositoryGRPC) (service grpcClient.Service) {
	connAuth := grpcAuthRepo.NewClient()

	impl := &implementation{
		pbAuth: auth.NewAuthenticationClient(connAuth),
	}

	return impl
}
