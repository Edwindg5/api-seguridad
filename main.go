//api-seguridad/main.go
package main

import (
	"api-seguridad/core/config"
	"api-seguridad/core/database"
	municipaldeps "api-seguridad/resources/municipalities/infrastructure/dependencies"
	municipalroutes "api-seguridad/resources/municipalities/infrastructure/routes"
	policedeps "api-seguridad/resources/police/infrastructure/dependencies"
	policeroutes "api-seguridad/resources/police/infrastructure/routes"
	requestdeps "api-seguridad/resources/request/infrastructure/dependencies"
	requestroutes "api-seguridad/resources/request/infrastructure/routes"
	roledeps "api-seguridad/resources/roles/infrastructure/dependencies"
	roleroutes "api-seguridad/resources/roles/infrastructure/routes"
	userdeps "api-seguridad/resources/users/infrastructure/dependencies"
	userroutes "api-seguridad/resources/users/infrastructure/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar base de datos
	database.InitDB()

	// Inicializar dependencias
	userdeps.InitDependencies()
	roledeps.InitDependencies()
	policedeps.InitDependencies()
	requestdeps.InitDependencies()
	municipaldeps.InitDependencies()

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