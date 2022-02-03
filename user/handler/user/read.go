package user

import (
	"context"
	"net/http"
	"user/handler/view"
	"user/service/user/input"

	"github.com/gin-gonic/gin"
)

// Read godoc
// @Tags Cinema
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get read of cinema
// @Description Return read of Cinema
// @Produce  json
// @Success 200 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /cinemas/:id [get]
func (ctrl *Controller) Read(c *gin.Context) {
	ctx := context.Background()

	input := &input.UserInput{
		ID: c.Param("id"),
	}

	item, err := ctrl.userService.Read(ctx, input)
	if err != nil {
		view.HandleError(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, item)
}
