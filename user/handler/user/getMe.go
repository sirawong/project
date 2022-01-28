package user

import (
	"context"
	"errors"
	"log"
	"net/http"
	"user/handler/view"
	"user/service/user/input"

	"github.com/gin-gonic/gin"
)

// GetMe godoc
// @Tags User
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get me of user
// @Description Return me of user
// @Produce  json
// @Success 200 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /users/me [get]
func (ctrl *Controller) GetMe(c *gin.Context) {
	ctx := context.Background()

	userId := c.GetString("userid")
	log.Println(userId)
	if userId == "" {
		view.HandleError(c.Writer, errors.New("connot get UserId"))
		return
	}

	input := &input.UserInput{
		ID: userId,
	}

	item, err := ctrl.userService.Read(ctx, input)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, item)
}
