package user

import (
	"context"
	"mime/multipart"
	"user/service/user/input"
	"user/service/user/output"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, in *input.UserInput) (out *output.User, token *string, err error)
	All(ctx context.Context) (items []*output.User, err error)
	Read(ctx context.Context, in *input.UserInput) (out *output.User, err error)
	Update(ctx context.Context, in *input.UserInput) (out *output.User, err error)
	Delete(ctx context.Context, in *input.UserInput) (err error)
	Upload(ctx context.Context, in *input.UserInput, filename string, file multipart.File) (out *output.User, err error)
}
