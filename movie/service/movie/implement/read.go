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

func (impl *implementation) Read(ctx context.Context, in *input.MovieInput) (out *output.Movie, err error) {
	item := &entities.Movie{}
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	err = impl.repo.Read(ctx, filters, item)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	return output.ParseToOutput(item), nil
}
