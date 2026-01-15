package api

import (
	"habit-tracker/internal/auth"
	"habit-tracker/internal/static"
	"net/http"
)

// Router devuelve el mux completo
func Router() http.Handler {
	mux := http.NewServeMux()

	// 1. Archivos estáticos (CSS, JS, imágenes…)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(static.FS())))

	// 2. Página de entrada
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "internal/static/web/index.html")
	})

	// 3. End-points de API
	mux.HandleFunc("/login", auth.LoginHandler)
	mux.HandleFunc("/register", auth.RegisterHandler)

	return mux
}
