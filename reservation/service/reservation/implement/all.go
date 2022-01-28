package implement

import (
	"context"
	"reservation/entities"
	"reservation/errs"
	"reservation/logs"
	"reservation/service/reservation/output"
)

func (impl *implementation) List(ctx context.Context, opt *entities.PageOption) (items []*output.Reservation, err error) {
	
	_, records, err := impl.repo.List(ctx, opt, &entities.Reservation{})
	if err != nil || len(records) == 0 {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	items = make([]*output.Reservation, len(records))
	for i, record := range records {
		items[i] = output.ParseToOutput(record.(*entities.Reservation))
	}

	return items, nil
}
