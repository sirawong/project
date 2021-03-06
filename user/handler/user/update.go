package user

import (
	"context"
	"errors"
	"net/http"
	"user/handler/view"
	"user/service/user/input"

	"github.com/gin-gonic/gin"
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
// @Router /users/:id [patch]
func (ctrl *Controller) Update(c *gin.Context) {
	ctx := context.Background()

	userId := c.GetString("userid")
	if userId == c.Param("id") {
		view.HandleError(c, errors.New("admin not allow to update yourself"))
		return
	}

	input := &input.UserInput{
		ID: c.Param("id"),
	}
	if err := c.ShouldBindJSON(input); err != nil {
		view.HandleError(c, err)
		return
	}

	item, err := ctrl.userService.Update(ctx, input)
	if err != nil {
		view.HandleError(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, item)
}
