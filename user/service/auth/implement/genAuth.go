package implement

import (
	"context"
	"errors"
	"fmt"
	"user/entities"
	"user/logs"
)

func (impl *implementation) GenAuth(ctx context.Context, ent *entities.User) (token *string, err error) {

	token, err = impl.jwt.GenerateToken(ent.ID)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("cannot generate token")
	}

	if token != nil {
		if ent.Tokens == nil {
			ent.Tokens = []*entities.Token{}
		}

		ent.Tokens = append(ent.Tokens, &entities.Token{Token: *token})

		filters := []string{
			fmt.Sprintf("_id:eq:%v", ent.ID),
		}

		err = impl.repo.Update(ctx, filters, ent)
		if err != nil {
			logs.Error(err)
			return nil, errors.New("cannot update token")
		}
	}

	return token, nil
}
