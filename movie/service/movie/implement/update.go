package implement

import (
	"context"
	"fmt"
	"movie/entities"
	"movie/errs"
	"movie/logs"
	"movie/service/movie/input"
	"movie/service/movie/output"
)

func (impl *implementation) Update(ctx context.Context, in *input.MovieInput) (out *output.Movie, err error) {

	ent := in.ParseToEntities()
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	movie := &entities.Movie{}
	err = impl.repo.Read(ctx, filters, movie)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	ent.Image = movie.Image

	err = impl.repo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	return output.ParseToOutput(ent), nil
}
