package implement

import (
	"context"
	"fmt"
	"log"
	"reservation/entities"
	"reservation/errs"
	"reservation/logs"
	"reservation/service/reservation/input"
	"reservation/service/reservation/output"
)

func (impl *implementation) Read(ctx context.Context, in *input.ReservationInput) (out *output.Reservation, err error) {
	cinema := &entities.Reservation{}
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	err = impl.repo.Read(ctx, filters, cinema)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}
	log.Println(cinema)
	return output.ParseToOutput(cinema), nil
}
