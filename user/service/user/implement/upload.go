package implement

import (
	"context"
	"fmt"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson"

	"user/entities"
	"user/errs"
	"user/logs"
	"user/service/user/input"
	"user/service/user/output"
)

func (impl *implementation) Upload(ctx context.Context, in *input.UserInput, filename string, file multipart.File) (out *output.User, err error) {
	urlPath, err := impl.storageRepo.Upload(ctx, filename, file)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	ent := bson.M{"imageurl": urlPath}
	err = impl.repo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	user := &entities.User{}
	err = impl.repo.Read(ctx, filters, user)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("User not found")
	}

	out = output.ParseToOutput(user)

	return out, nil
}
