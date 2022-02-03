package output

import (
	"cinema/entities"
)

type Cinema struct {
	ID             string  `json:"_id"`
	Name           string  `json:"name"`
	TicketPrice    float32 `json:"ticketPrice"`
	City           string  `json:"city"`
	Seats          [][]int `json:"seats"`
	SeatsAvailable int     `json:"seatsAvailable"`
	Image          string  `json:"image"`
}

func ParseToOutput(ent *entities.Cinema) (out *Cinema) {
	return &Cinema{
		ID:             ent.ID,
		Name:           ent.Name,
		TicketPrice:    ent.TicketPrice,
		City:           ent.City,
		Seats:          ent.Seats,
		SeatsAvailable: ent.SeatsAvailable,
		Image:          ent.Image,
	}
}
