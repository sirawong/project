package cinema

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	view "cinema/handler/view"
	"cinema/service/cinema/input"
)

// Update godoc
// @Tags User
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Update of User
// @Description Return Update of User
// @Produce  json
// @Success 200 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /cinemas/photo/:id [patch]
func (ctrl *Handlers) Upload(c *gin.Context) {
	ctx := context.Background()

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Please upload a file")
		return
	}

	defer file.Close()

	input := &input.CinemaInput{
		ID: c.Param("id"),
	}

	item, err := ctrl.service.Upload(ctx, input, fileHeader.Filename, file)
	if err != nil {
		view.HandleError(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, item)
}
