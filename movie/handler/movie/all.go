package movie

import (
	"context"

	"github.com/gin-gonic/gin"

	"movie/entities"
	view "movie/handler/view"
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
// @Router /movie [get]
func (ctrl *Handlers) All(c *gin.Context) {
	ctx := context.Background()

	opt := &entities.PageOption{}
	items, err := ctrl.service.List(ctx, opt)
	if err != nil {
		view.HandleError(c, err)
		return
	}

	view.MakeAllResp(c, items)
}
