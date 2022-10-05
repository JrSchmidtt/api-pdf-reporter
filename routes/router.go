package routes

import (
	"github.com/JrSchmidtt/api-gin/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		Generator := main.Group("pdf")
		{
			Generator.GET("/", controllers.CreatePDF)
		}
	}
	return router
}
