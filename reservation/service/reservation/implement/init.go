package implement

import (
	"reservation/config"
	repository "reservation/repository/mongodb"
	grpcService "reservation/service/grpcClient"
	"reservation/service/reservation"
	"reservation/utils"
)

type implementation struct {
	repo   repository.Repository
	uuid   utils.UUID
	grpc   grpcService.Service
	config *config.Config
}

func New(repo repository.Repository, uuid utils.UUID, grpc grpcService.Service, config *config.Config) (service reservation.Service) {
	return &implementation{repo: repo, uuid: uuid, grpc: grpc, config: config}
}
