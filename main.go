// api-seguridad/main.go
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

	"github.com/gin-contrib/cors"
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
	areachiefdeps.InitDependencies()
	requeststatusdeps.InitDependencies()
	requestdeps.InitDependencies()
	chiefs_periodsdeps.InitDependencies()
	request_detailsdeps.InitDependencies()
	permissionsdeps.InitDependencies()
	role_permissionsdeps.InitDependencies()

	// Crear router
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permitir cualquier origen
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 86400,
	}))

	// Configurar rutas base
	api := router.Group("/api/v1")
	{
		userroutes.ConfigureRoutes(api)
		roleroutes.ConfigureRoutes(api)
		policeroutes.ConfigureRoutes(api)
		municipalroutes.ConfigureRoutes(api)
		delegationroutes.ConfigureRoutes(api)
		typepoliceroutes.ConfigureRoutes(api)
		areachiefroutes.ConfigureRoutes(api)
		requeststatusroutes.ConfigureRoutes(api)
		requestroutes.ConfigureRoutes(api)
		chiefs_periodroutes.ConfigureRoutes(api)
		request_detailsroutes.ConfigureRoutes(api)
		permissionsroutes.ConfigureRoutes(api)
		role_permissionroutes.ConfigureRoutes(api)
	}

	// Iniciar servidor
	cfg := config.LoadConfig()
	log.Printf("Server running on port %s", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}