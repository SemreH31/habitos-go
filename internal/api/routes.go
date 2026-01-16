package api

import (
	"habit-tracker/internal/auth"
	"habit-tracker/internal/static"
	"io/fs" // Importante: añade este import
	"net/http"
	"strings"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	// 1. Obtenemos el embed.FS (que sí implementa fs.FS)
	fullFS := static.FS()

	// 2. Bajamos a la carpeta 'web/static' para los CSS/JS
	// Esto ya no dará error de compilación
	staticAssets, err := fs.Sub(fullFS, "web/static")
	if err != nil {
		panic(err)
	}

	// 3. Convertimos a http.Handler envolviendo con http.FS()
	fileServer := http.FileServer(http.FS(staticAssets))

	// Servimos archivos estáticos
	mux.Handle("/static/", http.StripPrefix("/static/", forceMimeType(fileServer)))

	// 4. Servir el index.html
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		// Leemos directamente del embed.FS
		index, err := fullFS.ReadFile("web/index.html")
		if err != nil {
			http.Error(w, "Archivo no encontrado", 404)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(index)
	})

	mux.HandleFunc("/login", auth.LoginHandler)
	mux.HandleFunc("/register", auth.RegisterHandler)

	return mux
}
func forceMimeType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if strings.HasSuffix(path, ".css") {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		} else if strings.HasSuffix(path, ".js") {
			// El estándar moderno es text/javascript o application/javascript
			w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
		}

		// Importante: Para ciberseguridad, esto previene el MIME Sniffing
		w.Header().Set("X-Content-Type-Options", "nosniff")

		next.ServeHTTP(w, r)
	})
}
