package cinema

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	view "cinema/handler/view"
	"cinema/service/cinema/input"
)

// Create godoc
// @Tags Cinema
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Create of cinema
// @Description Return Create of Cinema
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]output.Cinema}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /cinemas [post]
func (ctrl *Handlers) Create(c *gin.Context) {
	ctx := context.Background()

	input := &input.CinemaInput{}
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
