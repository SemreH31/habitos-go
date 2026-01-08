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
	// Tabla de usuarios
	usersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    );`
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

	_, err = DB.Exec(usersTable)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Base de datos y tablas listas.")
}
