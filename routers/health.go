package routers

import "github.com/gin-gonic/gin"

func init() {
	register("health", func(router *gin.Engine) {
		router.GET("/health", func(c *gin.Context) {
			c.JSON(200, "Hello ?")
		})
	})
}
