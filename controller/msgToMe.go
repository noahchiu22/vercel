package controller

import (
	"net/http"
	"vercel/models"
	"vercel/util"

	"github.com/gin-gonic/gin"
)

func MsgToMe(ctx *gin.Context) {
	g := util.Context{Ctx: ctx}
	var msg struct {
		Message string `json:"message"`
	}
	g.Ctx.ShouldBindJSON(&msg)

	util.Push(models.Push{
		To: util.TestUserId,
		Messages: []models.Message{{
			Type: "text",
			Text: msg.Message,
		}},
	}, util.TestToken)

	g.SendResponse(http.StatusOK, "ok", msg.Message)
}
