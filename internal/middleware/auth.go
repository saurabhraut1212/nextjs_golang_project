package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/auth"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/config"
)

func RequireAuth(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("access_token")
		if token == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}
		claims, err := auth.ParseToken(cfg.JWTSecret, token)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}
		c.Locals("userId", claims.UserID)
		return c.Next()
	}
}
