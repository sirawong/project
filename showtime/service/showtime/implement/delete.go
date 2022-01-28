package implement

import (
	"context"
	"fmt"
	"showtime/errs"
	"showtime/logs"
	"showtime/service/showtime/input"
)

func (impl *implementation) Delete(ctx context.Context, in *input.DeleteInput) (err error) {
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
