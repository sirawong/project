package reservation

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	view "reservation/handler/view"
	"reservation/service/reservation/input"
)

// Create godoc
// @Tags Suggest
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Create of cinema
// @Description Return Create of Cinema
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]output.Cinema}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /usermodeling/:username [post]
func (ctrl *Handlers) Suggest(c *gin.Context) {
	ctx := context.Background()

	input := &input.ReservationInput{
		Username: c.Param("username"),
	}

	data, err := ctrl.service.SuggestSeats(ctx, input)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, data)
}
