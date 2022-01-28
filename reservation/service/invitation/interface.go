package reservation

import (
	"context"
	"reservation/service/invitation/input"
)

//go:generate mockery --name=Service
type Service interface {
	Invitation(context.Context, []*input.InvitationInput) error
}
