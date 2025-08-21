package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/router"
)

func main() {
	// Read PORT from environment variable, fallback to 8000 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := router.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Server is running successfully on port %s!", port))
	})

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
