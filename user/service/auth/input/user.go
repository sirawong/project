package input

import (
	"user/entities"
)

type AuthInput struct {
	ID       string `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"-"`
}

func (input *AuthInput) ParseToEntities() (ent *entities.User) {
	ent = &entities.User{
		Username:   input.Username,
		Password:   input.Password,
	}

	return ent
}
