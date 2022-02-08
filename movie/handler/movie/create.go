package movie

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	view "movie/handler/view"
	"movie/service/movie/input"
)

// Create godoc
// @Tags Movie
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Create of Movie
// @Description Return Create of Movie
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]output.Movie}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /movies [post]
func (ctrl *Handlers) Create(c *gin.Context) {
	ctx := context.Background()

	input := &input.MovieInput{}
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
