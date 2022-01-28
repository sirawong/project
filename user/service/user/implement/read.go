package implement

import (
	"user/entities"
	"user/errs"
	"user/logs"
	"user/service/user/input"
	"user/service/user/output"

	"context"
	"fmt"
)

func (impl *implementation) Read(ctx context.Context, in *input.UserInput) (out *output.User, err error) {
	user := &entities.User{}
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	err = impl.repo.Read(ctx, filters, user)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	return output.ParseToOutput(user), nil
}
