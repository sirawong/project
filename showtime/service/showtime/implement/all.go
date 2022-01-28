package implement

import (
	"context"
	"showtime/entities"
	"showtime/errs"
	"showtime/logs"
	"showtime/service/showtime/output"
)

func (impl *implementation) All(ctx context.Context) (items []*output.ShowTime, err error) {
	opt := &entities.PageOption{}

	_, records, err := impl.repo.List(ctx, opt, &entities.ShowTime{})
	if err != nil || len(records) == 0 {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	items = make([]*output.ShowTime, len(records))
	for i, record := range records {
		items[i] = output.ParseToOutput(record.(*entities.ShowTime))
	}

	return items, nil
}
