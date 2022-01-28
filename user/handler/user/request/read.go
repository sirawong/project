package request

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"user/service/user/input"
)

func MakeReadReq(ctx *gin.Context, in *input.UserInput) (req *http.Request, err error) {

	ctx.Params = gin.Params{{
		Key: "id", Value: in.ID,
	}}

	req, err = http.NewRequest("GET", fmt.Sprintf("/users/%s", in.ID), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
