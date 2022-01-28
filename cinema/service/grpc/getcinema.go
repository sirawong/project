package grpc

import (
	"context"

	"cinema/logs"
	"cinema/service/cinema/input"
	pbCinema "cinema/service/grpc/protobuf/cinema"
)

func (impl *CinemaServer) GetCinemas(ctx context.Context, in *pbCinema.CinemaRequest) (resp *pbCinema.CinemaReply, err error) {
	data := make([]*pbCinema.Cinema, len(in.CinemaIds))

	for i := 0; i < len(in.CinemaIds); i++ {
		cinema, err := impl.cinemaSrv.Read(ctx, &input.CinemaInput{ID: in.CinemaIds[i]})
		if err != nil {
			logs.Error(err)
		}

		if err == nil {
			data[i] = &pbCinema.Cinema{
				CinemaId:       cinema.ID,
				Name:           cinema.Name,
				TicketPrice:    cinema.TicketPrice,
				City:           cinema.City,
				SeatsAvailable: cinema.SeatsAvailable,
				Image:          cinema.Image,
			}

			data[i].Seats = make([]*pbCinema.Seat, len(cinema.Seats))
			for j, seat := range cinema.Seats {
				var newSeat pbCinema.Seat
				newSeat.Seat = append(newSeat.Seat, seat...)
				data[i].Seats[j] = &newSeat
			}
		}
	}

	resp = &pbCinema.CinemaReply{}
	resp.Cinema = data

	return resp, nil
}
