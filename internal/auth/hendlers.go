package auth

import (
	"habit-tracker/internal/database"
	"log" // <--- Ahora sí lo usaremos
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "./web/index.html")
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		log.Printf("Intento de login para el correo: %s", email) // Log de auditoría

		var dbPassword string
		err := database.DB.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&dbPassword)

		if err != nil {
			log.Printf("Error: Usuario no encontrado o error en DB para %s: %v", email, err)
			http.Redirect(w, r, "/login?error=1", http.StatusSeeOther)
			return
		}

		if dbPassword != password {
			log.Printf("Advertencia: Contraseña incorrecta para el usuario: %s", email)
			http.Redirect(w, r, "/login?error=1", http.StatusSeeOther)
			return
		}

		log.Printf("Login exitoso: Usuario %s ha ingresado", email)

		// Crear cookie de sesión simple
		http.SetCookie(w, &http.Cookie{
			Name:  "session_token",
			Value: "usuario-autenticado",
			Path:  "/",
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}
