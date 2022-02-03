package implement

import (
	"context"
	"fmt"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson"

	"cinema/entities"
	"cinema/errs"
	"cinema/logs"
	"cinema/service/cinema/input"
	"cinema/service/cinema/output"
)

func (impl *implementation) Upload(ctx context.Context, in *input.CinemaInput, filename string, file multipart.File) (out *output.Cinema, err error){
	urlPath, err := impl.storageRepo.Upload(ctx, filename, file)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	ent := bson.M{"image": urlPath}
	err = impl.cinemaRepo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	cinema := &entities.Cinema{}
	err = impl.cinemaRepo.Read(ctx, filters, cinema)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("User not found")
	}

	out = output.ParseToOutput(cinema)

	return out, nil
}
