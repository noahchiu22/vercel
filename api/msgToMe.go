package handler

import (
	"net/http"
	"vercel/models"
	"vercel/util"

	"github.com/gin-gonic/gin"
)

func MsgToMe(g *gin.Context) {
	c := Context{g}
	var msg struct {
		Message string `json:"message"`
	}
	c.ShouldBindJSON(&msg)

	util.Push(models.Push{
		To: util.TestUserId,
		Messages: []models.Message{{
			Type: "text",
			Text: msg.Message,
		}},
	}, util.TestToken)

	c.SendResponse(http.StatusOK, "ok", msg.Message)
}
