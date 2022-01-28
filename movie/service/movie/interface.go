package movie

import (
	"bytes"
	"context"
	"mime/multipart"
	"movie/entities"
	"movie/service/movie/input"
	"movie/service/movie/output"
	"net/http"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, in *input.MovieInput) (out *output.Movie, err error)
	List(ctx context.Context, opt *entities.PageOption) (items []*output.Movie, err error)
	Read(ctx context.Context, in *input.MovieInput) (out *output.Movie, err error)
	Delete(ctx context.Context, in *input.MovieInput) (err error)
	Update(ctx context.Context, in *input.MovieInput) (out *output.Movie, err error)
	Upload(ctx context.Context, client HttpClienter, body *bytes.Buffer, writer *multipart.Writer, in *input.MovieInput) (out *output.Movie, err error)
}


//go:generate mockery --name=HttpClienter
type HttpClienter interface {
	Do(req *http.Request) (*http.Response, error)
}