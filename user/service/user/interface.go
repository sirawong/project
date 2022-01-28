package user

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
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
	Upload(ctx context.Context, client HttpClienter, body *bytes.Buffer, writer *multipart.Writer, in *input.UserInput) (out *output.User, err error)
}

//go:generate mockery --name=HttpClienter
type HttpClienter interface {
	Do(req *http.Request) (*http.Response, error)
}
