package movie

import (
	"context"
	"mime/multipart"
	"movie/entities"
	"movie/service/movie/input"
	"movie/service/movie/output"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, in *input.MovieInput) (out *output.Movie, err error)
	List(ctx context.Context, opt *entities.PageOption) (items []*output.Movie, err error)
	Read(ctx context.Context, in *input.MovieInput) (out *output.Movie, err error)
	Delete(ctx context.Context, in *input.MovieInput) (err error)
	Update(ctx context.Context, in *input.MovieInput) (out *output.Movie, err error)
	Upload(ctx context.Context, in *input.MovieInput, filename string, file multipart.File) (out *output.Movie, err error)
}
