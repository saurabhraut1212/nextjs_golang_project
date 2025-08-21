package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetAuthCookies(c *fiber.Ctx, access string, aExp time.Time, refresh string, rExp time.Time, prod bool) {
	secure := prod
	sameSite := fiber.CookieSameSiteLaxMode

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    access,
		HTTPOnly: true, Secure: secure, SameSite: sameSite,
		Expires: aExp, Path: "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refresh,
		HTTPOnly: true, Secure: secure, SameSite: sameSite,
		Expires: rExp, Path: "/",
	})
}

func ClearAuthCookies(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{Name: "access_token", Value: "", Expires: time.Unix(0, 0), Path: "/"})
	c.Cookie(&fiber.Cookie{Name: "refresh_token", Value: "", Expires: time.Unix(0, 0), Path: "/"})
}
