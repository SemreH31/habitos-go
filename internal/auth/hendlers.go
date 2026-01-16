package auth

import (
	"encoding/json"
	"habit-tracker/internal/database"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type registerReq struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req registerReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validaciones básicas
	if req.Password != req.PasswordConfirm {
		respJSON(w, http.StatusBadRequest, map[string]string{"error": "Passwords do not match"})
		return
	}
	if len(req.Password) < 6 {
		respJSON(w, http.StatusBadRequest, map[string]string{"error": "Password too short"})
		return
	}

	// Hasheo (cost=10 por defecto)
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		respJSON(w, http.StatusInternalServerError, map[string]string{"error": "Could not hash password"})
		return
	}

	// Insertamos
	id, err := database.CreateUser([]byte(req.Name), []byte(req.Email), hash)
	if err != nil {
		// ej. email único
		respJSON(w, http.StatusConflict, map[string]string{"error": "Email already registered"})
		return
	}

	respJSON(w, http.StatusCreated, map[string]any{"id": id, "message": "User created"})
}

// helper
func respJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req loginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		return
	}

	// 1. Buscamos el hash en BD
	hash, err := database.GetUserHashByEmail(req.Email)
	if err != nil {
		respJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		return
	}

	// 2. Comparamos hash
	if err := bcrypt.CompareHashAndPassword(hash, []byte(req.Password)); err != nil {
		respJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		return
	}

	// 3. Aquí generas JWT o cookie; ejemplo simple:
	respJSON(w, http.StatusOK, map[string]string{"message": "Login ok"})
}
