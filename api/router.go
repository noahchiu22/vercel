package handler

import (
	"vercel/middleware"

	"github.com/gin-gonic/gin"
)

func Setup_Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.Cors())
	g0 := router.Group("api")

	g0.POST("/msgToMe", MsgToMe)

	return router
}
