package grpcClient

import (
	auth "showtime/service/grpcClient/protobuf/auth"
)

type Service interface {
	VerifyToken(data *auth.TokenRequest) (res *auth.TokenReply, err error)
}
