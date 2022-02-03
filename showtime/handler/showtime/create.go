package showtime

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	view "showtime/handler/view"
	"showtime/service/showtime/input"
)

// Create godoc
// @Tags Showtime
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Create of showtime
// @Description Return Create of Showtime
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]output.Showtime}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /showtimes [post]
func (ctrl *Handlers) Create(c *gin.Context) {
	ctx := context.Background()

	input := &input.CreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.HandleError(c, err)
		return
	}

	items, err := ctrl.service.Create(ctx, input)
	if err != nil {
		view.HandleError(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusCreated, items)
}
