package implement

import (
	"context"
	"fmt"

	"cinema/errs"
	"cinema/logs"
	"cinema/service/cinema/input"
)

func (impl *implementation) Delete(ctx context.Context, in *input.CinemaInput) (err error) {
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}
	err = impl.cinemaRepo.Delete(ctx, filters)
	if err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	return nil
}
