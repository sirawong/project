package implement

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	auth "movie/service/grpcClient/protobuf/auth"
)

func (impl *implementation) VerifyToken(in *auth.TokenRequest) (*auth.TokenReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := impl.pbAuth.VerifyToken(ctx, in)
	if status.Code(err) != codes.OK {
		return nil, err
	}

	return res, nil
}
