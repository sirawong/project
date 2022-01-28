package entities

import "time"

type User struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Username  string    `bson:"username"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Role      string    `bson:"role"`
	Phone     string    `bson:"phone"`
	Imageurl  string    `bson:"imageurl"`
	Tokens    []*Token  `bson:"tokens"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type Token struct {
	Token string `bson:"token"`
}
