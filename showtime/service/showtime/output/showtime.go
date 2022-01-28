package output

import (
	"showtime/entities"
	"time"
)

type ShowTime struct {
	ID        string    `json:"_id"`
	StartAt   string    `json:"startAt"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	MovieId   string    `json:"movieId"`
	CinemaId  string    `json:"cinemaId"`
}

func ParseToOutput(ent *entities.ShowTime) (out *ShowTime) {
	return &ShowTime{
		ID:        ent.ID,
		StartAt:   ent.StartAt,
		StartDate: ent.StartDate,
		EndDate:   ent.EndDate,
		MovieId:   ent.MovieId,
		CinemaId:  ent.CinemaId,
	}
}
