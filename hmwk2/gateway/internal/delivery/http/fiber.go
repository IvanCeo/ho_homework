package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApp() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	return app
}
