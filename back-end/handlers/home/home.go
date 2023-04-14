package home

import "github.com/gofiber/fiber/v2"

func Test(c *fiber.Ctx) error {
	err := c.SendString("Api Is Up, Testing Route")
	return err
}