package main

import (
	"habit-tracker/internal/auth"
	"habit-tracker/internal/database"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	// 4. Servidor de archivos estáticos
	wd, _ := os.Getwd()
	staticDir := filepath.Join(wd, "web", "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// 1. Fijamos el header ANTES de cualquier escritura
		switch {
		case strings.HasSuffix(path, ".css"):
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		case strings.HasSuffix(path, ".js"):
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		default:
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		}
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// 2. Comprobamos que el archivo existe; si no, 404 claro
		file := filepath.Join(staticDir, path)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Println("Serving file:", file)
			http.NotFound(w, r)
			return
		}
		// 3. Ahora sí, servimos
		http.ServeFile(w, r, file) // mejor que fs.ServeHTTP para este caso
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
func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
