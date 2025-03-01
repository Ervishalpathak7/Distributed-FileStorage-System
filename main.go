package main

import (
	"context"
	"distributed-storage-system/services"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file")


func main() {

	// Initialize the Azure client
	_, err := services.InitAzureClient()
	if err != nil {
		panic(err)
	}

	// Connect to the Postgres database
	pgClient , err :=  services.InitPostgresClient()
	if err != nil {
		panic(err)
	}


	m , err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("PGUSER"),
			os.Getenv("PGPASSWORD"),
			os.Getenv("PGHOST"),
			os.Getenv("PGPORT"),	
			os.Getenv("PGDATABASE"),
			os.Getenv("PGSSLMODE"),
		),
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	

	// Run the migration
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("No migration was applied")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		log.Println("Migration successful")
	}


	// Set up the Routes 
	













	// Close the Postgres client
	defer pgClient.Close(context.Background())
}