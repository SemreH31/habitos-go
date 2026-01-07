package main

import (
	"habit-tracker/internal/database" // Ajusta esto al nombre de tu modulo
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Cargar el archivo .env
	err := godotenv.Load("../../.env") // Subimos dos niveles porque estamos en cmd/server/
	if err != nil {
		log.Println("No se encontró el archivo .env, usando variables del sistema")
	}

	// 2. Obtener variables de entorno
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./habits.db" // Valor por defecto
	}

	// 3. Inicializar Base de Datos
	log.Println("Iniciando conexión con la base de datos...")
	database.InitDB(dbPath)
	log.Println("Base de datos lista.")

	// 4. Mantener el programa vivo (por ahora)
	log.Println("El servidor de hábitos está configurado correctamente.")
}
