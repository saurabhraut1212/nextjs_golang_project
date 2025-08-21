package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/router"
)

func main() {
	app := router.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running successfully on port 8000!")
	})
	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}
}
