package models

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Hash  []byte `json:"-"` // nunca se serializa
}
