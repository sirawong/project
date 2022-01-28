package cinema

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
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

	fileMultipart, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Please upload a file")
		return
	}

	input := &input.CinemaInput{
		ID: c.Param("id"),
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", fileMultipart.Filename)
	file, _ := fileMultipart.Open()
	io.Copy(part, file)
	writer.Close()

	client := &http.Client{}

	item, err := ctrl.service.Upload(ctx, client, body, writer, input)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, item)
}
