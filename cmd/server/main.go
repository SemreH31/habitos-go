package main

import (
	"habit-tracker/internal/auth"
	"habit-tracker/internal/database"
	"log"
	"mime"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	mime.AddExtensionType(".css", "text/css; charset=utf-8")
	mime.AddExtensionType(".js", "application/javascript; charset=utf-8")

	// 1. Cargar el archivo .env
	err := godotenv.Load(".env") // Subimos dos niveles porque estamos en cmd/server/
	if err != nil {
		log.Println("No se encontró el archivo .env, usando variables del sistema")
	}

	// 2. Obtener variables de entorno
	dbPath := os.Getenv("DB_PATH")
	PORT := os.Getenv("PORT")
	if dbPath == "" {
		dbPath = "./habits.db" // Valor por defecto
	}
	if PORT == "" {
		PORT = "9000"
	}
	// 3. Inicializar Base de Datos
	log.Println("Iniciando conexión con la base de datos...")
	database.InitDB(dbPath)
	// Esto es solo para probar el login hoy
	query := `INSERT OR IGNORE INTO users (email, password) VALUES (?, ?)`
	_, _ = database.DB.Exec(query, "test@example.com", "123456")

	// 4. Servidor de archivos estáticos (Solución Nuclear)

	staticDir := "./web/static"
	fs := http.FileServer(http.Dir(staticDir))

	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// FORZADO ATÓMICO: Si falla aquí, no fallará en ningún lado
		if len(path) >= 4 && path[len(path)-4:] == ".css" {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		} else if len(path) >= 3 && path[len(path)-3:] == ".js" {
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		}

		w.Header().Set("X-Content-Type-Options", "nosniff")

		// LOG PARA DEBUG: Esto te dirá en la consola qué está pidiendo el navegador
		log.Printf("Sirviendo archivo estático: %s", path)

		fs.ServeHTTP(w, r)
	})))
	// 5. Rutas
	http.HandleFunc("/login", auth.LoginHandler)

	// 6. Iniciar Servidor
	log.Println("Servidor corriendo en http://localhost:9000/login")
	err = http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
