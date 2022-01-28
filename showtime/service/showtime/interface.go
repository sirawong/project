package showtime

import (
	"context"
	"showtime/service/showtime/input"
	"showtime/service/showtime/output"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, in *input.CreateInput) (out *output.ShowTime, err error)
	All(ctx context.Context) (out []*output.ShowTime, err error)
	Read(ctx context.Context, in *input.ReadInput) (out *output.ShowTime, err error)
	Delete(ctx context.Context, in *input.DeleteInput) (err error)
	Update(ctx context.Context, in *input.UpdateInput) (out *output.ShowTime, err error)
}