package grpcClient

import (
	auth "cinema/service/grpcClient/protobuf/auth"
)

//go:generate mockery --name=Service
type Service interface {
	VerifyToken(data *auth.TokenRequest) (res *auth.TokenReply, err error)
}
