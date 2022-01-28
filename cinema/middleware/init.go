package middleware

import grpcService "cinema/service/grpcClient"

type Service struct {
	GRPCSrv grpcService.Service
}

func New(GRPCSrv grpcService.Service) Service {
	return Service{GRPCSrv: GRPCSrv}
}
