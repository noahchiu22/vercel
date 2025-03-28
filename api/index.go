package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server := Setup_Router()

	server.ServeHTTP(w, r)
}

func (c *Context) SendResponse(status int, message string, data any) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}
