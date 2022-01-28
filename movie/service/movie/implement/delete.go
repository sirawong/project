package implement

import (
	"context"
	"fmt"
	"movie/errs"
	"movie/logs"
	"movie/service/movie/input"
)

func (impl *implementation) Delete(ctx context.Context, in *input.MovieInput) (err error) {
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
