package implement

import (
	"user/errs"
	"user/logs"
	"user/service/user/input"

	"context"
	"fmt"
)

func (impl *implementation) Delete(ctx context.Context, in *input.UserInput) (err error) {
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}
	err = impl.repo.Delete(ctx, filters)
	if err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	return nil
}
