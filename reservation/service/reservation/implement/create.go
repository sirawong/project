package implement

import (
	"context"
	"fmt"
	"reservation/errs"
	"reservation/logs"
	"reservation/service/reservation/input"
	"reservation/service/reservation/output"
	"reservation/utils"
)

func (impl *implementation) Create(ctx context.Context, in *input.ReservationInput) (out *output.Reservation, url string, err error) {
	in.ID = impl.uuid.Generate()

	url, err = utils.NewQRcode(fmt.Sprintf("%s/%s", impl.config.AppUrl, in.ID))
	if err != nil {
		logs.Error(err)
		errs.NewUnexpectedError()
	}

	ent := in.ParseToEntities()
	_, err = impl.repo.Create(ctx, ent)
	if err != nil {
		logs.Error(err)
		return nil, "", errs.NewBadRequestError(err.Error())
	}

	out = output.ParseToOutput(ent)

	return out, url, nil
}
