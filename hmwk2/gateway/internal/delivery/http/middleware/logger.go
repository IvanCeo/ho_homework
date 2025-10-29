package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Log() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		duration := time.Since((start))
		method := c.Method()
		path := c.Path()

		fmt.Printf("[%s] %s - %v\n", method, path, duration)

		return c.Next()
	}
}
