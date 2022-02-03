package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"user/handler/view"
	"user/service/auth/input"
)

// Login godoc
// @Tags User
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Login of User
// @Description Return Login of User
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]output.User}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /users/login [post]
func (ctrl *Controller) Login(c *gin.Context) {
	ctx := context.Background()

	input := &input.AuthInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.HandleError(c, err)
		return
	}

	user, token, err := ctrl.authService.Login(ctx, input)
	if err != nil {
		view.HandleError(c, err)
		return
	}

	data := make(map[string]interface{})

	data["user"] = user
	data["token"] = token

	view.MakeSuccessResp(c, http.StatusOK, data)
}
