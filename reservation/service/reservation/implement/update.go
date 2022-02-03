package implement

import (
	"context"
	"fmt"
	"reservation/entities"
	"reservation/errs"
	"reservation/logs"
	"reservation/service/reservation/input"
	"reservation/service/reservation/output"
)

func (impl *implementation) Update(ctx context.Context, in *input.ReservationInput) (out *output.Reservation, err error) {

	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}
	ent := in.ParseToEntities()

	err = impl.repo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	reserv := &entities.Reservation{}
	err = impl.repo.Read(ctx, filters, reserv)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	return output.ParseToOutput(reserv), nil
}
