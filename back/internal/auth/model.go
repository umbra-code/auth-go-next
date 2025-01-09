package auth

import "time"

type User struct {
	ID        uint   `json:"id" db:"id"`                 // Identificador único del usuario
	Username  string `json:"username" db:"username"`     // Nombre de usuario único
	Password  string `json:"password" db:"password"`     // Contraseña hasheada
	CreatedAt string `json:"created_at" db:"created_at"` // Fecha de creación del usuario
}

type RefreshToken struct {
	ID        uint      `json:"id" db:"id"`
	Token     string    `json:"token" db:"token"`
	UserID    uint      `json:"user_id" db:"user_id"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}
