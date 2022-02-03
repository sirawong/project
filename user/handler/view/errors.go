package view

import (
	"net/http"

	"user/errs"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case errs.AppError:
		c.AbortWithStatusJSON(e.Code, err)
	case error:
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
