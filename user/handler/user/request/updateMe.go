package request

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"user/service/user/input"
)

func MakeUpdateMeReq(ctx *gin.Context, in *input.UserInput) (req *http.Request, err error) {

	jsonBytes, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	ctx.Params = gin.Params{{
		Key: "id", Value: in.ID,
	}}

	req, err = http.NewRequest("PATCH", "/users/me", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func MakeUpdateMeReqInvalidJSON(in *input.UserInput) (req *http.Request, err error) {

	jsonBytes := []byte("{{{}}}")

	req, err = http.NewRequest("PATCH", "/users/me", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
