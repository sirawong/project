package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"user/handler/view"
	"user/service/auth/input"
)

// Logout godoc
// @Tags User
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Logout of User
// @Description Return Logout of User
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]output.User}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /users/logout [post]
func (ctrl *Controller) Logout(c *gin.Context) {
	ctx := context.Background()

	token, isTokenExist := c.Get("token")
	if !isTokenExist {
		view.HandleError(c, errors.New("connot get Token"))
	}

	userid, isUseridExist := c.Get("userid")
	if !isUseridExist {
		view.HandleError(c, errors.New("connot get UserId"))
	}

	input := &input.AuthInput{
		ID:    userid.(string),
		Token: token.(string),
	}

	err := ctrl.authService.Logout(ctx, input)
	if err != nil {
		view.HandleError(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusCreated, map[string]string{})
}
