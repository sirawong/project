package implement

import (
	"context"
	"fmt"
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

	err = impl.repo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	return output.ParseToOutput(ent), nil
}
