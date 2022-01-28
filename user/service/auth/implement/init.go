package implement

import (
	repository "user/repository/mongodb"
	"user/service/auth"
	"user/utils"
)

type implementation struct {
	repo repository.Repository
	jwt  utils.JWTService
}

func New(repo repository.Repository, jwt utils.JWTService) (service auth.Service) {
	return &implementation{repo: repo, jwt: jwt}
}
