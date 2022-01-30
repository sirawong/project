package implement

import (
	"context"
	"fmt"
	"log"
	"user/entities"
	"user/errs"
	"user/logs"
	"user/service/auth/input"
	"user/service/auth/output"

	"golang.org/x/crypto/bcrypt"
)

func (impl *implementation) Login(ctx context.Context, in *input.AuthInput) (out *output.User, token *string, err error) {
	user := &entities.User{}

	filters := []string{
		fmt.Sprintf("username:eq:%v", in.Username),
	}
	err = impl.repo.Read(ctx, filters, user)
	if user.ID == "" || err != nil {
		logs.Error(err)
		return nil, nil, errs.NewBadRequestError("You have entered an invalid username or password")
	}
	log.Println(user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		logs.Error(err)
		return nil, nil, errs.NewBadRequestError("You have entered an invalid username or password")
	}

	token, err = impl.GenAuth(ctx, user)
	if err != nil {
		return nil, nil, errs.NewBadRequestError(err.Error())
	}

	return output.ParseToOutput(user), token, nil
}
