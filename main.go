package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/lucasdasial/linkdash/internal/config"
	"github.com/lucasdasial/linkdash/internal/core/links"
	"github.com/lucasdasial/linkdash/internal/database"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(ctx, cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := echo.New()
	server.Use(middleware.RequestLogger())

	linksRepo := links.NewLinkRepo(db)
	linksHandler := links.NewLinkHandler(*linksRepo)

	linksHandler.RegisterRoutes(server.Group("/links"))
	server.Start(":" + cfg.Port)
}
