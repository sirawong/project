package reservation

import (
	"context"
	"reservation/entities"

	"github.com/gin-gonic/gin"

	view "reservation/handler/view"
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
func (ctrl *Handlers) All(c *gin.Context) {
	ctx := context.Background()

	opt := &entities.PageOption{}

	items, err := ctrl.service.List(ctx, opt)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	view.MakeAllResp(c, items)
}
