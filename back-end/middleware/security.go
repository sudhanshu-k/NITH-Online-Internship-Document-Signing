package middleware

import "github.com/gofiber/fiber/v2"

func Security(c *fiber.Ctx) error {
	c.Set("X-XSS-Protection", "0")
	c.Set("X-Frame-Options", "SAMEORIGIN")
	return c.Next()
}
