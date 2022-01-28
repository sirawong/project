package showtime

import (
	"context"

	"github.com/gin-gonic/gin"

	view "showtime/handler/view"
)

// All godoc
// @Tags Showtime
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get all of showtime
// @Description Return all of showtime
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=output.showtime}
// @Success 204 {object} view.ErrResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /showtimes [get]
func (ctrl *Handlers) All(c *gin.Context) {
	ctx := context.Background()

	items, err := ctrl.service.All(ctx)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	view.MakeAllResp(c, items)
}
