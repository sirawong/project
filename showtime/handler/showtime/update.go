package showtime

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	view "showtime/handler/view"
	"showtime/service/showtime/input"
)

// Update godoc
// @Tags Showtime
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Update of showtime
// @Description Return Update of Showtime
// @Produce  json
// @Success 200 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /showtimes/:id [patch]
func (ctrl *Handlers) Update(c *gin.Context) {
	ctx := context.Background()

	input := &input.UpdateInput{
		ID: c.Param("id"),
	}
	if err := c.ShouldBindJSON(input); err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	item, err := ctrl.service.Update(ctx, input)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, item)
}
