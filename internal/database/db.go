package database

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite" // Importación necesaria para el driver
)

var DB *sql.DB

func InitDB(filepath string) {
	var err error
	DB, err = sql.Open("sqlite", filepath)
	if err != nil {
		log.Fatal(err)
	}

	if DB == nil {
		log.Fatal("La base de datos es nula")
	}

	createTables()
}

func createTables() {
	// Tabla de hábitos
	habitsTable := `
	CREATE TABLE IF NOT EXISTS habits (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(habitsTable)
	if err != nil {
		log.Fatal("Error creando tabla habits:", err)
	}
}
