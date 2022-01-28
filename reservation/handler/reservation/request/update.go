package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"reservation/service/reservation/input"
)

func MakeUpdateReq(ctx *gin.Context, in *input.ReservationInput) (req *http.Request, err error) {

	jsonBytes, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	ctx.Params = gin.Params{{
		Key: "id", Value: in.ID,
	}}

	req, err = http.NewRequest("PATCH", fmt.Sprintf("/reservations/%s", in.ID), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func MakeUpdateReqInvalidJSON(in *input.ReservationInput) (req *http.Request, err error) {

	jsonBytes := []byte("{{{}}}")

	req, err = http.NewRequest("PATCH", fmt.Sprintf("/reservations/%s", in.ID), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
