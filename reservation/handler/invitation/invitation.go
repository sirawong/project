package invitation

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	view "reservation/handler/view"
	"reservation/service/invitation/input"
)

// Create godoc
// @Tags Invitation
// @Security BearerAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Summary Get Create of cinema
// @Description Return Create of Cinema
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]output.Cinema}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Router /cinemas [post]
func (ctrl *Handlers) Invitation(c *gin.Context) {
	ctx := context.Background()

	input := []*input.InvitationInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	err := ctrl.service.Invitation(ctx, input)
	if err != nil {
		view.HandleError(c.Writer, err)
		return
	}

	data := make(map[string]interface{}, 0)
	data["success"] = true
	data["msg"] = "The Invitation was sent!"

	view.MakeSuccessResp(c, http.StatusCreated, data)
}
