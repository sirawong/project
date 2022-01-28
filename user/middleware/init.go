package middleware

import (
	"user/service/auth"
)

type Service struct {
	auth auth.Service
}

func New(auth auth.Service) Service {
	return Service{auth: auth}
}
