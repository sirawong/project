package invitation

import (
	invitation "reservation/service/invitation"
)

type Handlers struct {
	service invitation.Service
}

func New(service invitation.Service) *Handlers {
	return &Handlers{service: service}
}
