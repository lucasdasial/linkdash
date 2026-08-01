package main

import (
	"context"
	"log"
	"lucasdasial/rupi/internal/database"
	httpserver "lucasdasial/rupi/internal/http_server"
)

func main() {

	pool, err := database.DbConnect(context.Background())
	if err != nil {

		log.Fatalf("Database connection error:\n %v", err)
	}
	defer pool.Close()

	log.Println("Database connected!")

	server := httpserver.New(":8080")

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
