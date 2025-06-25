package main

import (
	"log"
	"url_shortener/configs"
	"url_shortener/internal/url"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := configs.LoadConfig()

	dsn := cfg.MySQLDSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// Авто-миграция
	if err := db.AutoMigrate(&url.URL{}); err != nil {
		log.Fatalf("auto migration failed: %v", err)
	}

	app := fiber.New()

	repo := url.NewRepository(db)
	service := url.NewService(repo)
	handler := url.NewHandler(service)

	api := app.Group("/api/v1/url")
	api.Post("/shorten", handler.CreateShortURL)
	app.Get("/:id", handler.Redirect)

	log.Fatal(app.Listen(cfg.AppPort))
}
