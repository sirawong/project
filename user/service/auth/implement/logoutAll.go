package implement

import (
	"context"
	"fmt"
	"user/entities"
	"user/errs"
	"user/logs"
	"user/service/auth/input"
)

func (impl *implementation) LogoutAll(ctx context.Context, in *input.AuthInput) (err error) {
	user := &entities.User{}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	user.Tokens = []*entities.Token{}

	err = impl.repo.Update(ctx, filters, user)
	if err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	return nil
}
