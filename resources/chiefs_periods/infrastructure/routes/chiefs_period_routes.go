package routes

import (
	"api-seguridad/resources/chiefs_periods/infrastructure/controllers"
	"api-seguridad/resources/chiefs_periods/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureChiefsPeriodRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	// Initialize controllers
	createCtrl := controllers.NewCreateChiefsPeriodController(dependencies.GetCreateChiefsPeriodUseCase())
	getByIdCtrl := controllers.NewGetChiefsPeriodByIDController(dependencies.GetChiefsPeriodByIDUseCase())
	getAllCtrl := controllers.NewGetAllChiefsPeriodsController(dependencies.GetAllChiefsPeriodsUseCase())
	updateCtrl := controllers.NewUpdateChiefsPeriodController(dependencies.GetUpdateChiefsPeriodUseCase())
	deleteCtrl := controllers.NewSoftDeleteChiefsPeriodController(dependencies.GetSoftDeleteChiefsPeriodUseCase())
	getActiveCtrl := controllers.NewGetActiveChiefsPeriodController(dependencies.GetActiveChiefsPeriodUseCase())
	getByDateRangeCtrl := controllers.NewGetChiefsPeriodsByDateRangeController(dependencies.GetChiefsPeriodsByDateRangeUseCase())

	// Configure routes
	chiefsRoutes := router.Group("/chiefs-periods")
	{
		// CRUD endpoints
		chiefsRoutes.POST("", authMiddleware, createCtrl.Handle)
		chiefsRoutes.GET("", authMiddleware, getAllCtrl.Handle)
		chiefsRoutes.GET("/:id", authMiddleware, getByIdCtrl.Handle)
		chiefsRoutes.PUT("/:id", authMiddleware, updateCtrl.Handle)
		chiefsRoutes.DELETE("/:id", authMiddleware, deleteCtrl.Handle)

		// Special endpoints
		chiefsRoutes.GET("/active", authMiddleware, getActiveCtrl.Handle)
		chiefsRoutes.GET("/search", authMiddleware, getByDateRangeCtrl.Handle)
	}
}