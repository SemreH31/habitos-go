package main

import (
	"habit-tracker/internal/api"
	"habit-tracker/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./habits.db"
	}
	database.InitDB(dbPath)

	router := api.Router()

	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":9000"
	}
	log.Println("Server listening on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
