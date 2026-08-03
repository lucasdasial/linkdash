package main

import (
	"context"
	"log"

	"github.com/lucasdasial/linkdash/internal/config"
	databse "github.com/lucasdasial/linkdash/internal/database"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := databse.New(ctx, cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
