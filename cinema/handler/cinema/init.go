package cinema

import "cinema/service/cinema"

type Handlers struct {
	service cinema.CinemaService
}

func New(service cinema.CinemaService) *Handlers {
	return &Handlers{service: service}
}
