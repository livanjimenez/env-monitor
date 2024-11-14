package routers

import "github.com/gin-gonic/gin"

func RootRouters() {
	r := gin.Default()

	r.POST("/test", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is a test",
		})
	})
	
	r.Run(":8080")
}