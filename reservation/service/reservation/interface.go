package reservation

import (
	"context"
	"reservation/entities"
	"reservation/service/reservation/input"
	"reservation/service/reservation/output"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, in *input.ReservationInput) (out *output.Reservation, url string, err error)
	List(ctx context.Context, opt *entities.PageOption) (items []*output.Reservation, err error)
	Read(ctx context.Context, in *input.ReservationInput) (out *output.Reservation, err error)
	Delete(ctx context.Context, in *input.ReservationInput) (err error)
	Update(ctx context.Context, in *input.ReservationInput) (out *output.Reservation, err error)
	Checkin(ctx context.Context, in *input.ReservationInput) (out *output.Reservation, err error)
}
