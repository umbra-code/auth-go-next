package auth

import (
	"api-auth/pkg/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := AuthenticateUser(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generar Access Token
	accessToken, err := GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando token"})
		return
	}

	// Generar Refresh Token
	refreshToken, err := GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando refresh token"})
		return
	}

	// Guardar refresh token en la base de datos (opcional)
	err = SaveRefreshToken(user.ID, refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error guardando refresh token"})
		return
	}

	// Establecer cookies
	c.SetCookie("access-token", accessToken, 15*60, "/", "localhost", false, true)       // 15 minutos
	c.SetCookie("refresh-token", refreshToken, 7*24*3600, "/", "localhost", false, true) // 7 días

	c.JSON(http.StatusOK, gin.H{"message": "Login exitoso"})
}

func RegisterHandler(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := RegisterUser(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func UserProfileHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user ID"})
		return
	}

	// Obtener información del usuario desde la base de datos
	var username string
	err := database.DB.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "username": username})
}

func RefreshHandler(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh-token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token no encontrado"})
		return
	}

	// Validar el refresh token
	claims, err := ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generar un nuevo access token
	userID, _ := strconv.Atoi(claims.Subject)
	newAccessToken, err := GenerateAccessToken(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando nuevo access token"})
		return
	}

	// Establecer el nuevo access token en una cookie
	c.SetCookie("access-token", newAccessToken, 60*15, "/", "localhost", false, true) // 15 minutos

	c.JSON(http.StatusOK, gin.H{"message": "Token renovado con éxito"})
}
