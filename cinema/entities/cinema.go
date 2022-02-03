package entities

type Cinema struct {
	ID             string  `bson:"_id"`
	Name           string  `bson:"name"`
	TicketPrice    float32 `bson:"ticketPrice"`
	City           string  `bson:"city"`
	Seats          [][]int `bson:"seats"`
	SeatsAvailable int     `bson:"seatsAvailable"`
	Image          string  `bson:"image"`
}
