package auth

import (
	"context"
	"user/entities"
	"user/service/auth/input"
	"user/service/auth/output"
)

//go:generate mockery --name=Service
type Service interface {
	Login(ctx context.Context, in *input.AuthInput) (out *output.User, token *string, err error)
	Logout(ctx context.Context, in *input.AuthInput) (err error)
	LogoutAll(ctx context.Context, in *input.AuthInput) (err error)
	GenAuth(ctx context.Context, ent *entities.User) (token *string, err error)
	Verify(ctx context.Context, encodedToken string) (id string,  user *output.User, err error)
}
