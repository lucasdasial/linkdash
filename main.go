package main

import (
	"context"
	"log"
	"lucasdasial/rupi/internal/database"
)

func main() {
	logger := log.Default()

	pool, err := database.DbConnect(context.Background())
	if err != nil {

		logger.Fatalf("Error - database connection:\n %v", err)
	}
	defer pool.Close()

	logger.Println("Database connected!")

}
