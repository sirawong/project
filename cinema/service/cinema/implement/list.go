package implement

import (
	"context"

	"cinema/entities"
	"cinema/errs"
	"cinema/logs"
	"cinema/service/cinema/output"
)

func (impl *implementation) List(ctx context.Context, opt *entities.PageOption) (items []*output.Cinema, err error) {

	total, records, err := impl.cinemaRepo.List(ctx, opt, &entities.Cinema{})
	if err != nil || total == 0 {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	items = make([]*output.Cinema, len(records))
	for i, record := range records {
		items[i] = output.ParseToOutput(record.(*entities.Cinema))
	}

	return items, nil
}
