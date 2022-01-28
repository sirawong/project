package implement

import (
	repository "showtime/repository/mongodb"
	"showtime/service/showtime"
	"showtime/utils"
)

type implementation struct {
	repo repository.Repository
	uuid utils.UUID
}


func New(repo repository.Repository, uuid utils.UUID) (service showtime.Service) {
	return &implementation{repo: repo, uuid: uuid}
}
