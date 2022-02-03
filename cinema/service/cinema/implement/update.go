package implement

import (
	"context"
	"fmt"

	"cinema/entities"
	"cinema/errs"
	"cinema/logs"
	"cinema/service/cinema/input"
	"cinema/service/cinema/output"
)

func (impl *implementation) Update(ctx context.Context, in *input.CinemaInput) (out *output.Cinema, err error) {
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	cinema := &entities.Cinema{}
	err = impl.cinemaRepo.Read(ctx, filters, cinema)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}
	ent := in.ParseToEntities()
	ent.Image = cinema.Image

	err = impl.cinemaRepo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	return output.ParseToOutput(ent), nil
}
