package implement

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"user/entities"
	"user/errs"
	"user/logs"
	"user/service/user/input"
	"user/service/user/output"
)

func (impl *implementation) Update(ctx context.Context, in *input.UserInput) (out *output.User, err error) {
	filters := []string{
		fmt.Sprintf("_id:eq:%v", in.ID),
	}

	user := &entities.User{}
	err = impl.repo.Read(ctx, filters, user)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("User not found")
	}

	ent := in.ParseToEntities()

	if ent.Password != "" {
		hashing, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
		ent.Password = string(hashing)
	} else {
		ent.Password = user.Password
	}

	if user.Role == "" {
		ent.Role = user.Role
	}

	ent.Imageurl = user.Imageurl
	ent.Username = user.Username
	ent.CreatedAt = user.CreatedAt
	ent.UpdatedAt = time.Now()

	err = impl.repo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	return output.ParseToOutput(user), nil
}
