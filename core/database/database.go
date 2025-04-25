//api-seguridad/core/database/database.go
package database

import (
	"fmt"
	"log"

	"api-seguridad/core/config"

	entityareachiefs "api-seguridad/resources/area_chiefs/domain/entities"
	entitydelegations "api-seguridad/resources/delegation/domain/entities"
	entitymunicipalities "api-seguridad/resources/municipalities/domain/entities"
	entitypolice "api-seguridad/resources/police/domain/entities"
	entityrequest "api-seguridad/resources/request/domain/entities"
	entityrequeststatus "api-seguridad/resources/request_status/domain/entities"
	entityroles "api-seguridad/resources/roles/domain/entities"
	entitytypepolices "api-seguridad/resources/type_police/domain/entities"
	entityusers "api-seguridad/resources/users/domain/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction: true,
	})
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
	// Deshabilitar verificación de claves foráneas temporalmente
	if err := DB.Exec("SET FOREIGN_KEY_CHECKS=0").Error; err != nil {
		log.Printf("Warning: Could not disable foreign key checks: %v", err)
	}

	// Migrar tablas en orden adecuado
	tables := []interface{}{
		&entityroles.Role{},
		&entitytypepolices.TypePolice{},
		&entitymunicipalities.Municipality{},
		&entitydelegations.Delegation{},
		&entityareachiefs.AreaChief{},
		&entityusers.User{},
		&entitypolice.Police{},
		&entityrequeststatus.RequestStatus{},
		&entityrequest.Request{},
	}

	for _, table := range tables {
		if err := DB.AutoMigrate(table); err != nil {
			log.Printf("Warning: Error migrating table %T: %v", table, err)
		}
	}

	// Crear índices necesarios
	indexQueries := []string{
		"CREATE INDEX IF NOT EXISTS idx_users_role ON users(rol_id_fk)",
		"CREATE INDEX IF NOT EXISTS idx_users_creator ON users(created_by)",
		"CREATE INDEX IF NOT EXISTS idx_users_updater ON users(updated_by)",
	}

	for _, query := range indexQueries {
		if err := DB.Exec(query).Error; err != nil {
			log.Printf("Warning: Could not create index: %v", err)
		}
	}

	// Habilitar verificación de claves foráneas
	if err := DB.Exec("SET FOREIGN_KEY_CHECKS=1").Error; err != nil {
		log.Printf("Warning: Could not enable foreign key checks: %v", err)
	}

	log.Println("Migrations completed successfully")
}