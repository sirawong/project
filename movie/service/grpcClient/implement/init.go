package implement

import (
	"movie/repository/grpc"
	"movie/service/grpcClient"
	auth "movie/service/grpcClient/protobuf/auth"
)

type implementation struct {
	pbAuth auth.AuthenticationClient
}

func New(grpcAuthRepo grpc.RepositoryGRPC) (service grpcClient.Service) {
	connAuth, err := grpcAuthRepo.NewClient()
	if err != nil {
		return nil
	}

	impl := &implementation{
		pbAuth: auth.NewAuthenticationClient(connAuth),
	}

	return impl
}
