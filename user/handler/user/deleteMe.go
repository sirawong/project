package user

import (
	"context"
	"errors"
	"net/http"
	"user/handler/view"
	"user/service/user/input"

	"github.com/gin-gonic/gin"
)

// DeleteMe godoc
// @Tags User
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Delete of User
// @Description Return Delete of User
// @Produce  json
// @Success 200 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /users/me [delete]
func (ctrl *Controller) DeleteMe(c *gin.Context) {
	ctx := context.Background()

	role := c.GetString("role")
	if role == "" {
		view.HandleError(c.Writer, errors.New("connot get role"))
		return
	}

	if role != "superadmin" {
		view.HandleError(c.Writer, errors.New("you cannot delete yourself"))
		return
	}

	userid := c.GetString("userid")
	if userid == "" {
		view.HandleError(c.Writer, errors.New("connot get role"))
		return
	}

	input := &input.UserInput{
		ID: userid,
	}

	err := ctrl.userService.Delete(ctx, input)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	data := map[string]string{"message": "User Deleted"}

	view.MakeSuccessResp(c, http.StatusOK, data)
}
