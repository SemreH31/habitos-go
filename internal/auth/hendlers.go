package auth

import (
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// TODO: leer JSON, verificar contra BD, devolver JWT…
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "login ok"})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// TODO: insertar usuario
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "user created"})
}
