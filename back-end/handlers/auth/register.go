package auth

import (
	"context"

	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func Register(c *fiber.Ctx) error {
	db := database.DB
	json := new(model.User)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	if len(json.Password)<8{
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Length of password should be atleast 8.",
		})
	}
	json.Password = utils.HashAndSalt([]byte(json.Password))
	err := checkmail.ValidateFormat(json.Email)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Email Address",
		})
	}

	insertQuery := `insert into users values($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	_, err = db.Exec(context.Background(), insertQuery, uuid.New(), json.FirstName, json.LastName, json.Email, json.Password)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "User already exists",
		})
	}

	// fmt.Printf("res: %v\n", res)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}
