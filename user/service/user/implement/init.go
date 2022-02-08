package implement

import (
	"user/config"
	gStorage "user/repository/gstorage"
	repository "user/repository/mongodb"
	"user/service/auth"
	user "user/service/user"
	"user/utils"
)

type implementation struct {
	repo        repository.Repository
	auth        auth.Service
	uuid        utils.UUID
	config      *config.Config
	storageRepo gStorage.Storage
}

func New(repo repository.Repository, auth auth.Service, uuid utils.UUID, config *config.Config, storageRepo gStorage.Storage) (service user.Service) {
	return &implementation{repo: repo, auth: auth, uuid: uuid, config: config, storageRepo: storageRepo}
}
