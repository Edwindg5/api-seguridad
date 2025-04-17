package routes

import (

	"api-seguridad/resources/police/infrastructure/controllers"
	"api-seguridad/resources/police/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	policeController := controllers.NewPoliceController(dependencies.GetPoliceService())

	policeRoutes := router.Group("/police")
	{
		policeRoutes.POST("", policeController.CreatePolice)
		policeRoutes.GET("/:id", policeController.GetPolice)

	}
}