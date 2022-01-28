package implement

import (
	"user/entities"
	"user/errs"
	"user/logs"
	"user/service/user/output"

	"context"
)

func (impl *implementation) All(ctx context.Context) (items []*output.User, err error) {
	opt := &entities.PageOption{}

	_, records, err := impl.repo.List(ctx, opt, &entities.User{})
	if err != nil || len(records) == 0 {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	items = make([]*output.User, len(records))
	for i, record := range records {
		items[i] = output.ParseToOutput(record.(*entities.User))
	}

	return items, nil
}
