package entities

type Cinema struct {
	ID             string    `bson:"_id"`
	Name           string    `bson:"name"`
	TicketPrice    int32     `bson:"ticketPrice"`
	City           string    `bson:"city"`
	Seats          [][]int32 `bson:"seats"`
	SeatsAvailable int32     `bson:"seatsAvailable"`
	Image          string    `bson:"image"`
}
