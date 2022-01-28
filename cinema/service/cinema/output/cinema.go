package output

import (
	"cinema/entities"
)

type Cinema struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	TicketPrice    int32     `json:"ticketPrice"`
	City           string    `json:"city"`
	Seats          [][]int32 `json:"seats"`
	SeatsAvailable int32     `json:"seatsAvailable"`
	Image          string    `json:"image"`
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
