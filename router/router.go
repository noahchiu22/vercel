package router

import "github.com/gin-gonic/gin"

func Setup_Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router.Use(middleware.Cors())

	g0 := router.Group("api")

	g0.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	g0.GET("/ok", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	g0.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is post",
		})
	})

	return router
}
