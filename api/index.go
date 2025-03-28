package handler

import (
	"net/http"
	"vercel/controller"
	"vercel/middleware"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := Setup_Router()

	server.ServeHTTP(w, r)
}

func Setup_Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.Cors())
	g0 := router.Group("api")

	g0.POST("/msgToMe", controller.MsgToMe)

	return router
}
