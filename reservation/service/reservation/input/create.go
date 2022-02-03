package input

import (
	"reservation/entities"
	"time"
)

type ReservationInput struct {
	ID            string    `json:"-"`
	Date          time.Time `json:"date"`
	StartAt       string    `json:"startAt"`
	Seats         [][]int32 `json:"seats"`
	TicketPrice   int32     `json:"ticketPrice"`
	Total         int32     `json:"total"`
	ReservationId string    `json:"reservationId"`
	MovieId       string    `json:"movieId"`
	CinemaId      string    `json:"cinemaId"`
	Username      string    `json:"username"`
	Phone         string    `json:"phone"`
	Checkin       bool      `json:"checkin"`
}

func (input *ReservationInput) ParseToEntities() (ent *entities.Reservation) {
	return &entities.Reservation{
		ID:          input.ID,
		Date:        input.Date,
		StartAt:     input.StartAt,
		Seats:       input.Seats,
		TicketPrice: input.TicketPrice,
		Total:       input.Total,
		MovieId:     input.MovieId,
		CinemaId:    input.CinemaId,
		Username:    input.Username,
		Phone:       input.Phone,
		Checkin:     input.Checkin,
	}
}
