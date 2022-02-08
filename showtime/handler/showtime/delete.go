package showtime

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	view "showtime/handler/view"
	"showtime/service/showtime/input"
)

// Delete godoc
// @Tags Showtime
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Delete of showtime
// @Description Return Delete of Showtime
// @Produce json
// @Success 200 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /showtimes/:id [delete]
func (ctrl *Handlers) Delete(c *gin.Context) {
	ctx := context.Background()

	input := &input.DeleteInput{
		ID: c.Param("id"),
	}

	err := ctrl.service.Delete(ctx, input)
	if err != nil {
		view.HandleError(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, nil)
}
