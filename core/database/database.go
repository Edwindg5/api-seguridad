//api-seguridad/core/database/database.go
package database

import (
	"fmt"
	"log"

	"api-seguridad/core/config"

	entitydelegation "api-seguridad/resources/delegation/domain/entities"
	entitymunicipalities "api-seguridad/resources/municipalities/domain/entities"

	entitypolice "api-seguridad/resources/police/domain/entities"

	entityroles "api-seguridad/resources/roles/domain/entities"

	entitytypepolice "api-seguridad/resources/type_police/domain/entities"
	entityusers "api-seguridad/resources/users/domain/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
	RunMigrations()
}

func GetDB() *gorm.DB {
	return DB
}

func RunMigrations() {
    // 1. Primero tablas base sin dependencias
    err := DB.AutoMigrate(
        &entityroles.Role{},
        &entitytypepolice.TypePolice{},
        &entityusers.User{},
    )
    if err != nil {
        log.Fatalf("Failed to run base migrations: %v", err)
    }

    // 2. Luego municipalities que es referenciada
    err = DB.AutoMigrate(&entitymunicipalities.Municipality{})
    if err != nil {
        log.Fatalf("Failed to migrate Municipalities: %v", err)
    }

    // 3. Finalmente tablas con FKs
    err = DB.AutoMigrate(
        &entitydelegation.Delegation{},
        &entitypolice.Police{},
    )
    if err != nil {
        log.Fatalf("Failed to run dependent migrations: %v", err)
    }

    log.Println("All migrations completed successfully")
}