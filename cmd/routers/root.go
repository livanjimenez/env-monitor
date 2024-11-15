package routers

import (
	"github.com/gin-gonic/gin"
)

func RootRouters() {
	router := gin.Default()
	RegisterRoutes(router)
	router.Run(":8080")
}
