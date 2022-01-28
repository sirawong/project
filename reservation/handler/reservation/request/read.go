package request

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"reservation/service/reservation/input"
)

func MakeReadReq(ctx *gin.Context, in *input.ReservationInput) (req *http.Request, err error) {

	ctx.Params = gin.Params{{
		Key: "id", Value: in.ID,
	}}

	req, err = http.NewRequest("GET", fmt.Sprintf("/reservations/%s", in.ID), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
