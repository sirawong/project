package implement

import (
	"context"
	"showtime/errs"
	"showtime/logs"
	"showtime/service/showtime/input"
	"showtime/service/showtime/output"
)

func (impl *implementation) Create(ctx context.Context, in *input.CreateInput) (out *output.ShowTime, err error) {
	in.ID = impl.uuid.Generate()
	ent := in.ParseToEntities()
	_, err = impl.repo.Create(ctx, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	out = output.ParseToOutput(ent)

	return out, nil
}
