package main

import (
	"api-seguridad/core/config"
	"api-seguridad/core/database"
	municipalroutes "api-seguridad/resources/municipalities/infrastructure/routes"
	policeroutes "api-seguridad/resources/police/infrastructure/routes"
	requestroutes "api-seguridad/resources/request/infrastructure/routes"
	roleroutes "api-seguridad/resources/roles/infrastructure/routes"
	userroutes "api-seguridad/resources/users/infrastructure/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar base de datos
	database.InitDB()

	// Crear router
	router := gin.Default()

	// Configurar rutas base
	api := router.Group("/api/v1")
	{
		userroutes.ConfigureRoutes(api)
		roleroutes.ConfigureRoutes(api)
		policeroutes.ConfigureRoutes(api)
		requestroutes.ConfigureRoutes(api)
		municipalroutes.ConfigureRoutes(api)
	}

	// Iniciar servidor
	cfg := config.LoadConfig()
	log.Printf("Server running on port %s", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}