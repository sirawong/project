package implement

import (
	"context"
	"movie/entities"
	"movie/errs"
	"movie/logs"
	"movie/service/movie/output"
)

func (impl *implementation) List(ctx context.Context, opt *entities.PageOption) (items []*output.Movie, err error) {

	_, records, err := impl.repo.List(ctx, opt, &entities.Movie{})
	if err != nil || len(records) == 0 {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	items = make([]*output.Movie, len(records))
	for i, record := range records {
		items[i] = output.ParseToOutput(record.(*entities.Movie))
	}

	return items, nil
}
