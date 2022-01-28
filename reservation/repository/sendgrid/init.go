package sendgrid

import "reservation/entities"

type RepositorySendGrid interface {
	Send(data *entities.Email) (bool, error)
}

type SendGrid struct {
	sendgridApiKey string
}

func New(sendgridApiKey string) (sendGrid *SendGrid) {
	return &SendGrid{sendgridApiKey}
}
