package implement

import (
	"context"
	"fmt"
	"showtime/errs"
	"showtime/logs"
	"showtime/service/showtime/input"
	"showtime/service/showtime/output"
)

func (impl *implementation) Update(ctx context.Context, in *input.UpdateInput) (out *output.ShowTime, err error) {
	ent := in.ParseToEntities()
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}
	err = impl.repo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	return output.ParseToOutput(ent), nil
}
