package input

import (
	"cinema/entities"
)

type CinemaInput struct {
	ID             string  `json:"-"`
	Name           string  `json:"name"`
	TicketPrice    float32 `json:"ticketPrice"`
	City           string  `json:"city"`
	Seats          [][]int `json:"seats"`
	SeatsAvailable int     `json:"seatsAvailable"`
	Image          string  `json:"image"`
	CreatedBy      string  `json:"-"`
}

func (input *CinemaInput) ParseToEntities() (ent *entities.Cinema) {
	return &entities.Cinema{
		ID:             input.ID,
		Name:           input.Name,
		TicketPrice:    input.TicketPrice,
		City:           input.City,
		Seats:          input.Seats,
		SeatsAvailable: input.SeatsAvailable,
		Image:          input.Image,
	}
}
