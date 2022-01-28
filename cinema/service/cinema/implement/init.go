package implement

import (
	"cinema/config"
	repository "cinema/repository/mongodb"
	"cinema/service/cinema"
	"cinema/utils"
)

type implementation struct {
	cinemaRepo repository.CinemaRepository
	uuid       utils.UUID
	config     *config.Config
}

func New(cinemaRepo repository.CinemaRepository, uuid utils.UUID, appConfig *config.Config) (service cinema.CinemaService) {
	return &implementation{cinemaRepo: cinemaRepo, uuid: uuid, config: appConfig}
}
