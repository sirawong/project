package input

import (
	"showtime/entities"
	"time"
)

type CreateInput struct {
	ID         string    `json:"-"`
	StartAt    string    `json:"startAt"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	MovieId    string    `json:"movieId"`
	MovieName  string    `json:"movieName"`
	CinemaId   string    `json:"cinemaId"`
	CinemaName string    `json:"cinemaName"`
}

func (input *CreateInput) ParseToEntities() (ent *entities.ShowTime) {
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
