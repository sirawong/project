package cinema

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"

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
	Upload(ctx context.Context, client HttpClienter, body *bytes.Buffer, writer *multipart.Writer, in *input.CinemaInput) (out *output.Cinema, err error)
}

//go:generate mockery --name=HttpClienter
type HttpClienter interface {
	Do(req *http.Request) (*http.Response, error)
}
