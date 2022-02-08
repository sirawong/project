package entities

import "time"

type Movie struct {
	ID          string    `bson:"_id"`
	Title       string    `bson:"title"`
	Image       string    `bson:"image"`
	Language    string    `bson:"language"`
	Genre       string    `bson:"genre"`
	Director    string    `bson:"director"`
	Cast        string    `bson:"cast"`
	Description string    `bson:"description"`
	Duration    string    `bson:"duration"`
	ReleaseDate time.Time `bson:"releaseDate"`
	EndDate     time.Time `bson:"endDate"`
}
