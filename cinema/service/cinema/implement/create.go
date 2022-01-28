package implement

import (
	"context"

	"cinema/errs"
	"cinema/logs"
	"cinema/service/cinema/input"
	"cinema/service/cinema/output"
)

func (impl *implementation) Create(ctx context.Context, in *input.CinemaInput) (out *output.Cinema, err error) {
	in.ID = impl.uuid.Generate()
	ent := in.ParseToEntities()
	_, err = impl.cinemaRepo.Create(ctx, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	out = output.ParseToOutput(ent)

	return out, nil
}
