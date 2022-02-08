package entities

import "time"

type Reservation struct {
	ID          string    `bson:"_id"`
	Date        time.Time `bson:"date"`
	StartAt     string    `bson:"startAt"`
	Seats       [][]int32 `bson:"seats"`
	TicketPrice int32     `bson:"ticketPrice"`
	Total       int32     `bson:"total"`
	MovieId     string    `bson:"movieId"`
	CinemaId    string    `bson:"cinemaId"`
	Username    string    `bson:"username"`
	Phone       string    `bson:"phone"`
	Checkin     bool      `bson:"checkin"`
}
