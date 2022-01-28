package grpc

import (
	"context"
	"user/service/grpc/protobuf/auth"
)

func (impl *AuthenticationServer) VerifyToken(ctx context.Context, data *auth.TokenRequest) (resp *auth.TokenReply, err error) {
	resp = &auth.TokenReply{
		Id:     data.Id,
		Verify: false,
	}

	id, user, err := impl.authSrv.Verify(ctx, data.Token)
	if err != nil {
		return nil, err
	}

	if user != nil {
		resp.Role = user.Role
		resp.Id = id
		resp.Verify = true
	}

	return resp, nil
}
