package routes

import (
	"books-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(route *gin.Engine) *gin.Engine {
	main := route.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowBook)
		}
	}

	return route
}
