package entities

import "time"

type ShowTime struct {
	ID         string    `bson:"_id"`
	StartAt    string    `bson:"startAt"`
	StartDate  time.Time `bson:"startDate"`
	EndDate    time.Time `bson:"endDate"`
	MovieId    string    `bson:"movieId"`
	MovieName  string    `bson:"movieName"`
	CinemaId   string    `bson:"cinemaId"`
	CinemaName string    `bson:"cinemaName"`
}
