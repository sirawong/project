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
		return nil, errs.NewBadRequestError(err.Error())
	}

	ent := in.ParseToEntities()

	if ent.Password != "" {
		hashing, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
		ent.Password = string(hashing)
	}
	if ent.Role == "" {
		ent.Role = user.Role
	}

	ent.ID = user.ID
	ent.UpdatedAt = time.Now()

	err = impl.repo.Update(ctx, filters, ent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	return output.ParseToOutput(ent), nil
}
