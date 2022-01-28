package implement

import (
	"context"
	"reservation/logs"
	cinema "reservation/service/grpcClient/protobuf/cinema"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (impl *implementation) GetCinema(in *cinema.CinemaRequest) (data *cinema.CinemaReply, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := impl.pbCinema.GetCinemas(ctx, in)
	if status.Code(err) != codes.OK {
		logs.Error(err)
		return nil, err
	}

	return res, nil
}
