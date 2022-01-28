package reservation

import (
	"reservation/service/reservation"
)

type Handlers struct {
	service reservation.Service
}

func New(service reservation.Service) *Handlers {
	return &Handlers{service: service}
}
