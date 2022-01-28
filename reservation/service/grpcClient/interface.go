package grpcClient

import (
	auth "reservation/service/grpcClient/protobuf/auth"
	cinema "reservation/service/grpcClient/protobuf/cinema"
)

//go:generate mockery --name=Service
type Service interface {
	VerifyToken(data *auth.TokenRequest) (res *auth.TokenReply, err error)
	GetCinema(in *cinema.CinemaRequest) (data *cinema.CinemaReply, err error)
}
