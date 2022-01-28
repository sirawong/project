package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"user/handler/view"
	"user/service/auth/input"
)

// LogoutAll godoc
// @Tags User
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get LogoutAll of User
// @Description Return LogoutAll of User
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]output.User}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /users/logoutAll [post]
func (ctrl *Controller) LogoutAll(c *gin.Context) {
	ctx := context.Background()

	input := &input.AuthInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	err := ctrl.authService.LogoutAll(ctx, input)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusCreated, nil)
}
