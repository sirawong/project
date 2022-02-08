package implement

import (
	"context"
	"fmt"
	"reservation/errs"
	"reservation/logs"
	"reservation/service/reservation/input"

	"go.mongodb.org/mongo-driver/bson"
)

func (impl *implementation) Checkin(ctx context.Context, in *input.ReservationInput) (err error) {

	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	ent := bson.M{"checkin": true}
	err = impl.repo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return errs.NewBadRequestError(err.Error())
	}

	return nil
}
