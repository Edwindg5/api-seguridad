// api-seguridad/resources/area_chiefs/infrastructure/routes/area_chief_routes.go)
package routes

import (
	"api-seguridad/resources/area_chiefs/infrastructure/controllers"
	"api-seguridad/resources/area_chiefs/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	createCtrl := controllers.NewCreateAreaChiefController(dependencies.GetCreateAreaChiefUseCase())
	getByIDCtrl := controllers.NewGetAreaChiefByIDController(dependencies.GetAreaChiefByIDUseCase())
	getAllCtrl := controllers.NewGetAllAreaChiefsController(dependencies.GetAllAreaChiefsUseCase())
	updateCtrl := controllers.NewUpdateAreaChiefController(dependencies.GetUpdateAreaChiefUseCase())
	deleteCtrl := controllers.NewDeleteAreaChiefController(dependencies.GetDeleteAreaChiefUseCase())

	// Define routes
	chiefRoutes := router.Group("/area-chiefs")
	{
		chiefRoutes.POST("", createCtrl.Handle)          // POST /area-chiefs
		chiefRoutes.GET("", getAllCtrl.Handle)          // GET /area-chiefs (list all)
		chiefRoutes.GET("/:id", getByIDCtrl.Handle)     // GET /area-chiefs/:id
		chiefRoutes.PUT("/:id", updateCtrl.Handle)      // PUT /area-chiefs/:id
		chiefRoutes.DELETE("/:id", deleteCtrl.Handle)   // DELETE /area-chiefs/:id (soft delete)
	}
}