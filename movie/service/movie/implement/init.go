package implement

import (
	"movie/config"
	gStorage "movie/repository/gstorage"
	repository "movie/repository/mongodb"
	"movie/service/movie"
	"movie/utils"
)

type implementation struct {
	repo        repository.Repository
	uuid        utils.UUID
	config      *config.Config
	storageRepo gStorage.Storage
}

func New(repo repository.Repository, uuid utils.UUID, appConfig *config.Config, storageRepo gStorage.Storage) (service movie.Service) {
	return &implementation{repo: repo, uuid: uuid, config: appConfig, storageRepo: storageRepo}
}
