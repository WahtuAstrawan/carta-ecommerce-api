package main

import (
	"log"

	"carta-ecommerce-api/internal/config"
	"carta-ecommerce-api/internal/infrastructure/db"
	"carta-ecommerce-api/internal/infrastructure/db/migration"
)

func main() {
	dbConfig := config.LoadDatabaseConfig()

	database, err := db.NewPostgres(dbConfig)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	if err := migration.Run(database, dbConfig); err != nil {
		log.Fatal("migration failed:", err)
	}

	log.Println("database connected and migration applied")
}
