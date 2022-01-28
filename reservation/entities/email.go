package entities

type Email struct {
	From         string `json:"from"`
	FromEmail    string `json:"from_email"`
	To           string `json:"to"`
	ToEmail      string `json:"to_email"`
	Subject      string `json:"subject"`
	ContentPlain string `json:"content_plain"`
	ContentHtml  string `json:"content_html"`
}
