package main

import (
	"github.com/Wilo/url-shortener-echo/application/controllers"
	"github.com/Wilo/url-shortener-echo/domain/services"
	"github.com/Wilo/url-shortener-echo/infrastructure/persistence"
	"github.com/labstack/echo/v4"
)

func main() {
	// PG connection
	pg, err := persistence.NewPostgresConnection()
	if err != nil {
		panic(err)
	}

	// Migrations
	err = persistence.NewPostgresMigrations(pg).Migrate()
	if err != nil {
		panic(err)
	}

	// Infraestructure Layer
	urlRepo := persistence.NewUrlPostgresRepo(pg)

	// Domain Layer

	// Services
	urlService := services.NewUrlServices(urlRepo)

	// Application Layer
	// Init Echo
	server := echo.New()

	//Init Controllers
	urlController := controllers.NewUrlController(urlService)
	urlController.InitUrlControllerRouting(server)

	server.Logger.Fatal(server.Start("127.0.0.1:1323"))
}
