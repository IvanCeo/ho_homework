package main

import (
	"log"
	"github.com/gofiber/fiber/v3"
)

func pingHandler(c fiber.Ctx) error {
	return c.SendString("pong")
}

func main() {
	app := fiber.New()

	app.Get("/ping", pingHandler)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("server err: %v", err)
		}
	}()

	select{}
}