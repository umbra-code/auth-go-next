package database

import (
	"api-auth/config"
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Cambia el driver
)

var DB *sql.DB

func InitDatabase() {
	var err error

	// Configurar SQLite
	if config.DBDriver == "sqlite" {
		DB, err = sql.Open("sqlite", config.DBName)
		if err != nil {
			log.Fatalf("Error al conectar a la base de datos SQLite: %v", err)
		}
		log.Println("Conexión exitosa a SQLite")
	} else {
		log.Fatalf("Base de datos no soportada: %s", config.DBDriver)
	}

	// Prueba de conexión
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error al verificar la conexión: %v", err)
	}

	// Crear tablas si no existen
	createTables()
}

func createTables() {
	// Crear tabla de usuarios
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		log.Fatalf("Error al crear la tabla de usuarios: %v", err)
	}

	log.Println("Tabla de usuarios creada o ya existe")

	// Crear tabla de refresh tokens
	createRefreshTable := `
	CREATE TABLE IF NOT EXISTS refresh_tokens (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		token TEXT NOT NULL UNIQUE,
		expires_at DATETIME NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err = DB.Exec(createRefreshTable)
	if err != nil {
		log.Fatalf("Error al crear la tabla de refresh tokens: %v", err)
	}

	log.Println("Tabla de refresh tokens creada o ya existe")
}
