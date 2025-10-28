package main

import (
	"log"

	"gateway/internal/delivery/http"
	"gateway/internal/delivery/http/ledger"
	"gateway/internal/delivery/http/router"
	"ledger/pkg/factory"
)

// func pingHandler(c fiber.Ctx) error {
// 	return c.SendString("pong")
// }

func main() {
	server := factory.NewLedgerFactory()
	handler := ledger.NewHandler(server)

	app := http.NewApp()

	router.Route(app, handler)

	log.Fatal(app.Listen(":8080"))
}
