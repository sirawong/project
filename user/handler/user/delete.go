package user

import (
	"context"
	"net/http"
	"user/handler/view"
	"user/service/user/input"

	"github.com/gin-gonic/gin"
)

// Delete godoc
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
// @Router /users/:id [delete]
func (ctrl *Controller) Delete(c *gin.Context) {
	ctx := context.Background()

	input := &input.UserInput{
		ID: c.Param("id"),
	}

	err := ctrl.userService.Delete(ctx, input)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	data := map[string]string{"message": "User Deleted"}

	view.MakeSuccessResp(c, http.StatusOK, data)
}
