package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SecretToken, DBDriver, DBName string

func LoadConfig() {
	// Carga las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar .env: %v", err)
	}

	// Leer variables específicas
	SecretToken = getEnv("SECRET_TOKEN", "default_secret")
	DBDriver = getEnv("DB_DRIVER", "sqlite")
	DBName = getEnv("DB_NAME", "./data.db")
}

// Función auxiliar para obtener variables con un valor predeterminado
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
