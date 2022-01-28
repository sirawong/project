package implement

import (
	"context"
	"fmt"
	"log"

	"cinema/entities"
	"cinema/errs"
	"cinema/logs"
	"cinema/service/cinema/input"
	"cinema/service/cinema/output"
)

func (impl *implementation) Read(ctx context.Context, in *input.CinemaInput) (out *output.Cinema, err error) {
	cinema := &entities.Cinema{}
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	err = impl.cinemaRepo.Read(ctx, filters, cinema)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}
	log.Println(cinema)
	return output.ParseToOutput(cinema), nil
}
