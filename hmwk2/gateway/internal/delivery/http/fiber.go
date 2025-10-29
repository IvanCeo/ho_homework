package http

import (
	"gateway/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApp() *fiber.App {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(middleware.Log())
	app.Use(middleware.SetHeaders())
	// app.Use(logger.New(logger.Config{
	// 	Format: "[${method}] ${path} - ${latency}\n",
	// }))

	return app
}
