package cinema

import (
	"context"
	"mime/multipart"

	"cinema/entities"
	"cinema/service/cinema/input"
	"cinema/service/cinema/output"
)

//go:generate mockery --name=CinemaService
type CinemaService interface {
	Create(ctx context.Context, in *input.CinemaInput) (out *output.Cinema, err error)
	List(ctx context.Context, opt *entities.PageOption) (items []*output.Cinema, err error)
	Read(ctx context.Context, in *input.CinemaInput) (out *output.Cinema, err error)
	Delete(ctx context.Context, in *input.CinemaInput) (err error)
	Update(ctx context.Context, in *input.CinemaInput) (out *output.Cinema, err error)
	Upload(ctx context.Context, in *input.CinemaInput, filename string, file multipart.File) (out *output.Cinema, err error)
}