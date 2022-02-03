package implement

import (
	"context"
	"fmt"
	"reservation/entities"
	"reservation/errs"
	"reservation/logs"
	"reservation/service/reservation/input"
	"reservation/service/reservation/output"

	"go.mongodb.org/mongo-driver/bson"
)

func (impl *implementation) Checkin(ctx context.Context, in *input.ReservationInput) (out *output.Reservation, err error) {

	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	ent := bson.M{"checkin": true}
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
