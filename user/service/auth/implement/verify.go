package implement

import (
	"context"
	"errors"
	"fmt"
	"user/entities"
	"user/errs"
	"user/logs"
	"user/service/auth/output"
)

func (impl *implementation) Verify(ctx context.Context, encodedToken string) (id string, user *output.User, err error) {

	id, err = impl.jwt.ValidateToken(encodedToken)
	if err != nil {
		logs.Error(err)
		return "", nil, errs.NewUnauthorizedError()
	}

	userEnt := &entities.User{}
	filter := []string{
		fmt.Sprintf("tokens.token:eq:%v", encodedToken),
		fmt.Sprintf("_id:eq:%v", id),
	}
	err = impl.repo.Read(ctx, filter, userEnt)
	if err != nil {
		logs.Error(err)
		return "", nil, errs.NewUnexpectedError()
	}

	if userEnt.Role == "" {
		logs.Error(errors.New("role not found"))
		return "", nil, errs.NewUnauthorizedError()
	}

	return id, output.ParseToOutput(userEnt), nil
}
