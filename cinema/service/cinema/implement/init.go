package implement

import (
	"cinema/config"
	gStorage "cinema/repository/gstorage"
	repository "cinema/repository/mongodb"
	"cinema/service/cinema"
	"cinema/utils"
)

type implementation struct {
	cinemaRepo repository.CinemaRepository
	uuid       utils.UUID
	config     *config.Config
	storageRepo gStorage.Storage
}

func New(cinemaRepo repository.CinemaRepository, uuid utils.UUID, appConfig *config.Config, storageRepo gStorage.Storage) (service cinema.CinemaService) {
	return &implementation{cinemaRepo: cinemaRepo, uuid: uuid, config: appConfig, storageRepo:storageRepo}
}
