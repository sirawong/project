package request

import (
	"fmt"
	"net/http"
	"user/service/user/input"

	"github.com/gin-gonic/gin"
)

func MakeDeleteReq(ctx *gin.Context, in *input.UserInput) (req *http.Request, err error) {

	ctx.Params = gin.Params{{
		Key: "id", Value: in.ID,
	}}

	req, err = http.NewRequest("DELETE", fmt.Sprintf("/users/%s", in.ID), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
