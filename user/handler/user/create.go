package user

import (
	"context"
	"net/http"
	"user/handler/view"
	"user/service/user/input"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Tags User
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Create of User
// @Description Return Create of User
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]output.User}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /users [post]
func (ctrl *Controller) Create(c *gin.Context) {
	ctx := context.Background()

	input := &input.UserInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.HandleError(c, err)
		return
	}

	user, token, err := ctrl.userService.Create(ctx, input)
	if err != nil {
		view.HandleError(c, err)
		return
	}

	data := make(map[string]interface{})

	data["user"] = user
	data["token"] = token

	view.MakeSuccessResp(c, http.StatusCreated, data)
}
