package home

import "github.com/gofiber/fiber/v2"
import "time"

func Test(c *fiber.Ctx) error {
	err := c.SendString("Api Is Up, Testing Route")
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "sdf",
		Path:     "/",
		Secure:   false,
		Expires:   time.Now().Add(60 * time.Minute),
		HTTPOnly: true,
	})
	return err
}
