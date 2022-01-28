package showtime

import (
	"showtime/service/showtime"
)

type Handlers struct {
	service showtime.Service
}

func New(service showtime.Service) *Handlers {
	return &Handlers{service: service}
}
