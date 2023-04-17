package profile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(model.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
