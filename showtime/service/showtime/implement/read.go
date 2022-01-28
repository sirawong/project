package implement

import (
	"context"
	"fmt"
	"log"
	"showtime/entities"
	"showtime/errs"
	"showtime/logs"
	"showtime/service/showtime/input"
	"showtime/service/showtime/output"
)

func (impl *implementation) Read(ctx context.Context, in *input.ReadInput) (out *output.ShowTime, err error) {
	showtime := &entities.ShowTime{}
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	err = impl.repo.Read(ctx, filters, showtime)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}
	log.Println(showtime)
	return output.ParseToOutput(showtime), nil
}
