package output

type SuggestSeats struct {
	Positions       map[string]int `json:"positions"`
	NumberOfTickets int            `json:"numberOfTickets"`
}
