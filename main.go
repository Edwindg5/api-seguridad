//api-seguridad/main.go
package main

import (
	"api-seguridad/core/config"
	"api-seguridad/core/database"
	areachiefdeps "api-seguridad/resources/area_chiefs/infrastructure/dependencies"
	areachiefroutes "api-seguridad/resources/area_chiefs/infrastructure/routes"
	delegationdeps "api-seguridad/resources/delegation/infrastructure/dependencies"
	delegationroutes "api-seguridad/resources/delegation/infrastructure/routes"
	municipaldeps "api-seguridad/resources/municipalities/infrastructure/dependencies"
	municipalroutes "api-seguridad/resources/municipalities/infrastructure/routes"
	policedeps "api-seguridad/resources/police/infrastructure/dependencies"
	policeroutes "api-seguridad/resources/police/infrastructure/routes"
	requestdeps "api-seguridad/resources/request/infrastructure/dependencies"
	requestroutes "api-seguridad/resources/request/infrastructure/routes"
	requeststatusdeps "api-seguridad/resources/request_status/infrastructure/dependencies"
	requeststatusroutes "api-seguridad/resources/request_status/infrastructure/routes"
	roledeps "api-seguridad/resources/roles/infrastructure/dependencies"
	roleroutes "api-seguridad/resources/roles/infrastructure/routes"
	userdeps "api-seguridad/resources/users/infrastructure/dependencies"
	userroutes "api-seguridad/resources/users/infrastructure/routes"
	typepolicedeps "api-seguridad/resources/type_police/infrastructure/dependencies"
	typepoliceroutes "api-seguridad/resources/type_police/infrastructure/routes"
	chiefs_periodsdeps "api-seguridad/resources/chiefs_periods/infrastructure/dependencies"
	chiefs_periodroutes "api-seguridad/resources/chiefs_periods/infrastructure/routes"
	request_detailsdeps "api-seguridad/resources/request_details/infrastructure/dependencies"
	request_detailsroutes "api-seguridad/resources/request_details/infrastructure/routes"
	permissionsdeps "api-seguridad/resources/permissions/infrastructure/dependencies"
	permissionsroutes "api-seguridad/resources/permissions/infrastructure/routes"
	role_permissionsdeps "api-seguridad/resources/role_permissions/infrastructure/dependencies"
	role_permissionroutes "api-seguridad/resources/role_permissions/infrastructure/routes"
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
	municipaldeps.InitDependencies()
	delegationdeps.InitDependencies()
	typepolicedeps.InitDependencies()
	areachiefdeps.InitDependencies() // Nueva dependencia
	requeststatusdeps.InitDependencies() // Nueva dependencia
	requestdeps.InitDependencies() // Nueva dependencia
	chiefs_periodsdeps.InitDependencies() // Nueva dependencia
	request_detailsdeps.InitDependencies() // Nueva dependencia
	permissionsdeps.InitDependencies() // Nueva dependencia
	role_permissionsdeps.InitDependencies() // Nueva dependencia

	// Crear router
	router := gin.Default()

	// Configurar rutas base
	api := router.Group("/api/v1")
	{
		userroutes.ConfigureRoutes(api)
		roleroutes.ConfigureRoutes(api)
		policeroutes.ConfigureRoutes(api)
		municipalroutes.ConfigureRoutes(api)
		delegationroutes.ConfigureRoutes(api)
		typepoliceroutes.ConfigureRoutes(api)
		areachiefroutes.ConfigureRoutes(api) // Nueva ruta
		requeststatusroutes.ConfigureRoutes(api) // Nueva ruta
		requestroutes.ConfigureRoutes(api) // Nueva ruta
		chiefs_periodroutes.ConfigureRoutes(api) // Nueva ruta
		request_detailsroutes.ConfigureRoutes(api) // Nueva ruta
		permissionsroutes.ConfigureRoutes(api) // Nueva ruta
		role_permissionroutes.ConfigureRoutes(api) // Nueva ruta
	}

	// Iniciar servidor
	cfg := config.LoadConfig()
	log.Printf("Server running on port %s", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}