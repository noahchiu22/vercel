package router

import "github.com/gin-gonic/gin"

func Setup_Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router.Use(middleware.Cors())

	return router
}
