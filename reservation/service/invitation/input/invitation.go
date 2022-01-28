package input

type InvitationInput struct {
	To     string `json:"to"`
	Host   string `json:"host"`
	Movie  string `json:"movie"`
	Time   string `json:"time"`
	Date   string `json:"date"`
	Cinema string `json:"cinema"`
	Image  string `json:"image"`
	Seat   string `json:"seat"`
}
