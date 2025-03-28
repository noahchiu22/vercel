package util

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	Ctx *gin.Context
}

func (c *Context) SendResponse(status int, message string, data any) {
	c.Ctx.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}
