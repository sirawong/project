package input

import (
	"showtime/entities"
	"time"
)

type UpdateInput struct {
	ID         string    `json:"_id"`
	StartAt    string    `json:"startAt"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	MovieId    string    `json:"movieId"`
	MovieName  string    `json:"movieName"`
	CinemaId   string    `json:"cinemaId"`
	CinemaName string    `json:"cinemaName"`
}

func (input *UpdateInput) ParseToEntities() (ent *entities.ShowTime) {
	return &entities.ShowTime{
		ID:         input.ID,
		StartAt:    input.StartAt,
		StartDate:  input.StartDate,
		EndDate:    input.EndDate,
		MovieId:    input.MovieId,
		MovieName:  input.MovieName,
		CinemaId:   input.CinemaId,
		CinemaName: input.CinemaName,
	}
}
