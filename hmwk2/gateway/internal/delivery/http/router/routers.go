package router

import (
	"gateway/internal/delivery/http/ledger"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, handler *ledger.Handler) {
	api := app.Group("/api")

	api.Post("/transactions", (handler.CreateTransactionHandle))
	api.Get("/transactions", handler.GetTransactionsHandle)
	api.Post("/budgets", handler.CreateBudgetHandle)
	api.Get("/budgets", handler.GetBudgetsHandle)
}
