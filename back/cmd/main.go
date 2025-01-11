package main

import (
	"api-auth/config"
	"api-auth/pkg/database"
	"api-auth/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuraciones y conectar la base de datos
	config.LoadConfig()
	database.InitDatabase()

	// Configurar servidor
	router := gin.Default()
	// Configurar las reglas de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Orígenes permitidos
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true, // Permitir cookies o credenciales
		MaxAge:           12 * time.Hour,
	}))
	routes.SetupRoutes(router)

	// Arrancar servidor
	log.Println("Servidor escuchando en http://localhost:8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
