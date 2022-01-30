package user

import (
	"context"
	"errors"
	"net/http"
	"user/handler/view"
	"user/service/user/input"

	"github.com/gin-gonic/gin"
)

// UpdateMe godoc
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
// @Router /users/me [patch]
func (ctrl *Controller) UpdateMe(c *gin.Context) {
	ctx := context.Background()

	userId := c.GetString("userid")
	if userId == "" {
		view.HandleError(c.Writer, errors.New("connot get UserId"))
		return
	}

	role := c.GetString("role")
	if role == "superadmin" {
		view.HandleError(c.Writer, errors.New("admin not allow to update yourself"))
		return
	}

	input := &input.UserInput{
		ID: userId,
	}
	if err := c.ShouldBindJSON(input); err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	item, err := ctrl.userService.Update(ctx, input)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, item)
}
