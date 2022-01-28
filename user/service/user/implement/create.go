package implement

import (
	"context"
	"time"
	"user/errs"
	"user/logs"
	"user/service/user/input"
	"user/service/user/output"

	"golang.org/x/crypto/bcrypt"
)

func (impl *implementation) Create(ctx context.Context, in *input.UserInput) (out *output.User, token *string, err error) {
	in.ID = impl.uuid.Generate()
	if in.Role == "" {
		in.Role = "guest"
	}

	hashing, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		logs.Error(err)
		return nil, nil, errs.NewUnexpectedError()
	}
	password := string(hashing)
	in.Password = password

	ent := in.ParseToEntities()

	ent.CreatedAt = time.Now()
	ent.UpdatedAt = time.Now()

	_, err = impl.repo.Create(ctx, ent)
	if err != nil {
		logs.Error(err)
		return nil, nil, errs.NewBadRequestError(err.Error())
	}

	token, err = impl.auth.GenAuth(ctx, ent)
	if err != nil {
		return nil, nil, errs.NewBadRequestError(err.Error())
	}

	out = output.ParseToOutput(ent)

	return out, token, nil
}
