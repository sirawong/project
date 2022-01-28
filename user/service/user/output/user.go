package output

import (
	"time"
	"user/entities"
)

type User struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Phone     string    `json:"phone"`
	Imageurl  string    `json:"imageurl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ParseToOutput(ent *entities.User) (out *User) {
	out = &User{
		ID:        ent.ID,
		Name:      ent.Name,
		Username:  ent.Username,
		Email:     ent.Email,
		Role:      ent.Role,
		Phone:     ent.Phone,
		Imageurl:  ent.Imageurl,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}

	return out
}
