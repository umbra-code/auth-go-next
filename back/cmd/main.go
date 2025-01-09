package main

import (
	"api-auth/config"
	"api-auth/pkg/database"
	"api-auth/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuraciones y conectar la base de datos
	config.LoadConfig()
	database.InitDatabase()

	// Configurar servidor
	router := gin.Default()
	routes.SetupRoutes(router)

	// Arrancar servidor
	log.Println("Servidor escuchando en http://localhost:8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
