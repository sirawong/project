package user

import "user/service/user"

type Controller struct {
	userService user.Service
}

func New(service user.Service) *Controller {
	return &Controller{userService: service}
}
