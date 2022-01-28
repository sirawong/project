package implement

import (
	"movie/config"
	repository "movie/repository/mongodb"
	"movie/service/movie"
	"movie/utils"
)

type implementation struct {
	repo   repository.Repository
	uuid   utils.UUID
	config *config.Config
}

func New(repo repository.Repository, uuid utils.UUID, appConfig *config.Config) (service movie.Service) {
	return &implementation{repo: repo, uuid: uuid, config: appConfig}
}
