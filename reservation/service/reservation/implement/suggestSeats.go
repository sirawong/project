package implement

import (
	"context"
	"fmt"
	"log"
	"math"
	"reservation/entities"
	"reservation/errs"
	"reservation/logs"
	pbCinema "reservation/service/grpcClient/protobuf/cinema"
	"reservation/service/reservation/input"
	"reservation/service/reservation/output"
)

func (impl *implementation) SuggestSeats(ctx context.Context, in *input.ReservationInput) (out *output.SuggestSeats, err error) {

	opt := &entities.PageOption{}
	opt.Filters = append(opt.Filters, fmt.Sprintf("username:eq:%v", in.Username))
	total, userReservations, err := impl.repo.List(ctx, opt, &entities.Reservation{})
	if err != nil || total == 0 {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	CinemaId := make([]string, len(userReservations))
	for i, reservate := range userReservations {
		CinemaId[i] = reservate.(*entities.Reservation).CinemaId
	}
	
	data := &pbCinema.CinemaRequest{
		CinemaIds: CinemaId,
	}
	cinemas, err := impl.grpc.GetCinema(data)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}


	numberOfSeats := 0
	numberOfTicketsArray := 0
	positions := map[string]int{}

	for _, reservate := range userReservations {
		reservation := reservate.(*entities.Reservation)
		for _, cinema := range cinemas.Cinema {
			if reservation.CinemaId == cinema.CinemaId {
				//find how many rows the cinema has
				position := impl.getPosition(len(cinema.Seats), reservation.Seats)
				positions[position] += 1
				numberOfTicketsArray++
				numberOfSeats += len(reservation.Seats)
			}
		}
	}
	
	
	numberOfTickets := int(math.Round(float64(numberOfSeats / numberOfTicketsArray)))

	out = &output.SuggestSeats{
		Positions:       positions,
		NumberOfTickets: numberOfTickets,
	}

	return out, nil
}

func (impl *implementation) getPosition(cinemaRows int, seats [][]int32) string {
	var rowNum = int(seats[0][0])
	var step = cinemaRows / 3
	var pos = 1

	log.Println("Test")

	for i := step; i <= cinemaRows; i += step {
		if rowNum < i {
			switch pos {
			case 1:
				return "front"
			case 2:
				return "center"
			case 3:
				return "back"
			}
		}
		pos++
	}
	return ""
}
