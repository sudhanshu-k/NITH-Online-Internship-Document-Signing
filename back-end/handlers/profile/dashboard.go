package profile

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func Dashboard(c *fiber.Ctx) error {
	user := c.Locals("user").(model.UserResponse)

	if !user.IsFaculty {
		getAllForms := "select idform1 from user_to_form where iduser=$1"

		rows, err := database.DB.Query(context.Background(), getAllForms, user.ID.String())
		// utils.FatalError(rows.Err())

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "DB Error."})
		}

		if !rows.Next() {
			return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{"status": "success", "message": "No forms filled.", "rows": ""})
		}

		var data uuid.UUID
		rows.Scan(&data)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Filled Form", "rows": rows})
	}

	

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "rows": "jh", "filled": 1})
}
