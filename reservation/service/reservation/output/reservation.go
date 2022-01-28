package output

import (
	"reservation/entities"
	"time"
)

type Reservation struct {
	ID            string    `json:"_id"`
	Date          time.Time `json:"date"`
	StartAt       string    `json:"startAt"`
	Seats         [][]int32 `json:"seats"`
	TicketPrice   int32     `json:"ticketPrice"`
	Total         int32     `json:"total"`
	CinemaId      string    `json:"cinemaId"`
	Username      string    `json:"username"`
	Phone         string    `json:"phone"`
	Checkin       bool      `json:"checkin"`
}

func ParseToOutput(ent *entities.Reservation) (out *Reservation) {
	return &Reservation{
		ID:            ent.ID,
		Date:          ent.Date,
		StartAt:       ent.StartAt,
		Seats:         ent.Seats,
		TicketPrice:   ent.TicketPrice,
		Total:         ent.Total,
		CinemaId:      ent.CinemaId,
		Username:      ent.Username,
		Phone:         ent.Phone,
		Checkin:       ent.Checkin,
	}
}
