package implement

import (
	"context"
	"fmt"
	"user/entities"
	"user/errs"
	"user/logs"
	"user/service/auth/input"
)

func (impl *implementation) Logout(ctx context.Context, in *input.AuthInput) (err error) {
	user := &entities.User{}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
		fmt.Sprintf("tokens.token:eq:%v", in.Token),
	}

	err = impl.repo.Read(ctx, filters, user)
	if err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	for i, token := range user.Tokens {
		if token.Token == in.Token {
			user.Tokens = append(user.Tokens[:i], user.Tokens[i+1:]...)
		}
	}

	err = impl.repo.Update(ctx, filters, user)
	if err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	return nil
}
