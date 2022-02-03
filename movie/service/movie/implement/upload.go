package implement

import (
	"context"
	"fmt"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson"

	"movie/entities"
	"movie/errs"
	"movie/logs"
	"movie/service/movie/input"
	"movie/service/movie/output"
)

func (impl *implementation) Upload(ctx context.Context, in *input.MovieInput, filename string, file multipart.File) (out *output.Movie, err error) {
	urlPath, err := impl.storageRepo.Upload(ctx, filename, file)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	ent := bson.M{"image": urlPath}
	err = impl.repo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	user := &entities.Movie{}
	err = impl.repo.Read(ctx, filters, user)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("User not found")
	}

	out = output.ParseToOutput(user)

	return out, nil
}
