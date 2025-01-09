package auth

import (
	"api-auth/config"
	"api-auth/pkg/database"
	"api-auth/pkg/utils"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims personalizados
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateRefreshToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 7 días

	// Crear los claims
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   fmt.Sprint(userID),
		},
	}

	// Crear el token firmado
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(config.SecretToken)

	// Firmar el token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Generar un token JWT
func GenerateAccessToken(userID uint) (string, error) {
	// Tiempo de expiración del token (15 minutos)
	expirationTime := time.Now().Add(15 * time.Minute)

	// Crear los claims
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   fmt.Sprint(userID),
		},
	}

	// Crear el token firmado
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(config.SecretToken)

	// Firmar el token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Validar un token JWT
func ValidateToken(tokenString string) (*Claims, error) {
	secret := []byte(config.SecretToken)

	// Parsear el token y verificarlo
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Verificar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extraer los claims del token
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func AuthenticateUser(username, password string) (User, error) {
	var hashedPassword sql.NullString
	var userID sql.NullInt64
	var createdAt sql.NullTime

	// Consulta de la base de datos
	err := database.DB.QueryRow("SELECT id, password, created_at FROM users WHERE username = ?", username).Scan(&userID, &hashedPassword, &createdAt)
	if err == sql.ErrNoRows {
		return User{}, errors.New("invalid username or password")
	} else if err != nil {
		return User{}, err
	}

	// Validar contraseña
	if !utils.CheckPasswordHash(password, hashedPassword.String) {
		return User{}, errors.New("invalid username or password")
	}

	return User{ID: uint(userID.Int64), Username: username, CreatedAt: createdAt.Time.String()}, nil
}

func RegisterUser(username, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
	return err
}

func SaveRefreshToken(userID uint, token string) error {
	query := `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES (?, ?, ?)`
	_, err := database.DB.Exec(query, userID, token, time.Now().Add(7*24*time.Hour))
	return err
}

func ValidateRefreshToken(userID uint, token string) (bool, error) {
	query := `SELECT COUNT(*) FROM refresh_tokens WHERE user_id = ? AND token = ?`
	var count int
	err := database.DB.QueryRow(query, userID, token).Scan(&count)
	return count > 0, err
}
