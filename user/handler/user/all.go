package user

import (
	"context"
	"user/handler/view"

	"github.com/gin-gonic/gin"
)

// All godoc
// @Tags Cinema
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get all of cinema
// @Description Return all of cinema
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=output.cinema}
// @Success 204 {object} view.ErrResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /cinemas [get]
func (ctrl *Controller) All(c *gin.Context) {
	ctx := context.Background()

	items, err := ctrl.userService.All(ctx)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	view.MakeAllResp(c, items)
}
