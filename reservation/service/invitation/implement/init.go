package implement

import (
	repository "reservation/repository/sendgrid"
	invitation "reservation/service/invitation"
)

type implementation struct {
	SendGridRepo repository.RepositorySendGrid
}

func New(SendGridRepo repository.RepositorySendGrid) (service invitation.Service) {
	return &implementation{SendGridRepo: SendGridRepo}
}
