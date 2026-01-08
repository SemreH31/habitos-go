package auth

import (
	"database/sql"
	"fmt"
	"habit-tracker/internal/database"
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

		var dbPassword string
		// Buscamos al usuario por email
		err := database.DB.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&dbPassword)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Fprint(w, "Usuario no encontrado")
			} else {
				fmt.Fprint(w, "Error en el servidor")
			}
			return
		}

		// Comparamos la contraseña (por ahora en texto plano para probar)
		if password == dbPassword {
			fmt.Fprint(w, "¡Bienvenido! Login exitoso.")
		} else {
			fmt.Fprint(w, "Contraseña incorrecta")
		}
	}
}
