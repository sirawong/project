package implement

import (
	"context"
	"movie/errs"
	"movie/logs"
	"movie/service/movie/input"
	"movie/service/movie/output"
)

func (impl *implementation) Create(ctx context.Context, in *input.MovieInput) (out *output.Movie, err error) {
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
