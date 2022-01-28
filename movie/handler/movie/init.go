package movie

import (
	"movie/service/movie"
)

type Handlers struct {
	service movie.Service
}

func New(service movie.Service) *Handlers {
	return &Handlers{service: service}
}
