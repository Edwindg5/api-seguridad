package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error al cargar .env: %v", err)
	}

	// Obtener variables de entorno en el mismo orden
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// Crear string de conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	// Conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Error al abrir conexión: %v", err)
	}
	defer db.Close()

	// Verificar si la conexión es exitosa
	if err = db.Ping(); err != nil {
		log.Fatalf("❌ No se pudo conectar a la base de datos: %v", err)
	}

	log.Println("✅ Conexión exitosa a la base de datos 🎉")
}
