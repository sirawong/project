package auth

import "user/service/auth"

type Controller struct {
	authService auth.Service
}

func New(service auth.Service) *Controller {
	return &Controller{authService: service}
}
