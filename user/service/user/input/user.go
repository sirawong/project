package input

import (
	"user/entities"
)

type UserInput struct {
	ID       string `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Imageurl string `json:"imageurl"`
	Token    string `json:"-"`
}

func (input *UserInput) ParseToEntities() (ent *entities.User) {
	ent = &entities.User{
		ID:       input.ID,
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
		Role:     input.Role,
		Imageurl: input.Imageurl,
	}

	return ent
}
