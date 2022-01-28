package request

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"movie/service/movie/input"
)

func MakeDeleteReq(ctx *gin.Context, in *input.MovieInput) (req *http.Request, err error) {

	ctx.Params = gin.Params{{
		Key: "id", Value: in.ID,
	}}

	req, err = http.NewRequest("DELETE", fmt.Sprintf("/movies/%s", in.ID), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
