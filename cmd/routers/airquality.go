package routers

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	airQuality := r.Group("/airquality")
	{
		airQuality.GET("/", GetAirQuality)
		airQuality.POST("/", PostAirQuality)
		airQuality.PUT("/", PutAirQuality)
		airQuality.DELETE("/", DeleteAirQuality)
	}
}

func GetAirQuality(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAirQuality",
	})
}

func PostAirQuality(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PostAirQuality",
	})
}

func PutAirQuality(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PutAirQuality",
	})
}

func DeleteAirQuality(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteAirQuality",
	})
}
