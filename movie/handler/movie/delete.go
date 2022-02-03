package movie

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	view "movie/handler/view"
	"movie/service/movie/input"
)

// Delete godoc
// @Tags Cinema
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Delete of cinema
// @Description Return Delete of Cinema
// @Produce  json
// @Success 200 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /cinemas/:id [delete]
func (ctrl *Handlers) Delete(c *gin.Context) {
	ctx := context.Background()

	input := &input.MovieInput{
		ID: c.Param("id"),
	}

	err := ctrl.service.Delete(ctx, input)
	if err != nil {
		view.HandleError(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, nil)
}
