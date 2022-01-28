package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeSuccessResp(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func MakeAllResp(c *gin.Context, items interface{}) {
	MakeSuccessResp(c, http.StatusOK, items)
}
