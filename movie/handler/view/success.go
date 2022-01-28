package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeSuccessResp(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func MakeAllResp(c *gin.Context, items interface{}) {
	// c.Header(xContentLength, strconv.Itoa(total))
	MakeSuccessResp(c, http.StatusOK, items)
}
