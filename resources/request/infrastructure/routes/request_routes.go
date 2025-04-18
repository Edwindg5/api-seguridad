//api-seguridad/resources/request/infrastructure/routes/request_routes.go
package routes

import (

	"api-seguridad/resources/request/infrastructure/controllers"
	"api-seguridad/resources/request/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	requestController := controllers.NewRequestController(dependencies.GetRequestService())

	requestRoutes := router.Group("/requests")
	{
		requestRoutes.POST("", requestController.CreateRequest)
		requestRoutes.GET("/:id", requestController.GetRequest)
		
	}
}