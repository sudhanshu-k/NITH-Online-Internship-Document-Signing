package form

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	// "github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/google/uuid"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
)

func PostUgIntern(c *fiber.Ctx) error {
	user := c.Locals("user").(model.UserResponse)

	getAllForms := "select idform1 from user_to_form where iduser=$1"

	rows, _ := database.DB.Query(context.Background(), getAllForms, user.ID.String())
	if rows.Err() != nil {
		utils.Logger.Error("Database query execution resulted in error.", zap.Error(rows.Err()))

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    404,
			"message": "Server Error",
		})
	}

	if rows.Next() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Form already filled."})
	}

	var formData model.UGInternForm
	if err := c.BodyParser(&formData); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	uuidForm := uuid.New()
	insertQuery := `insert into uginternform values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, 
		$15, $16, $17, $18, $19)`
	res, err := database.DB.Exec(context.TODO(), insertQuery, formData.Name, formData.FatherName, formData.Address, formData.Contact,
		formData.CompanyName, formData.AreaOfIntrest, formData.IsOffline, formData.StartDay, formData.EndDay,
		formData.Weeks, formData.FromTPO, formData.Stipend, formData.FormDate, formData.RemarksDept, formData.RemarksFI,
		uuidForm, time.Now(), time.Now(), formData.Email)

	if err != nil {
		utils.Logger.Error("Database query execution resulted in error.", zap.Error(err))

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    404,
			"message": "Server Error",
		})
	}

	fmt.Printf("res: %v\n", res)
	fmt.Printf("user.ID: %v\n", user.ID.String())
	fmt.Printf("user.ID: %v\n", uuidForm.String())
	insertQuery = "insert into user_to_form values($1, $2)"
	_, err = database.DB.Exec(context.TODO(), insertQuery, user.ID.String(), uuidForm.String())

	if err != nil {
		utils.Logger.Error("Database query execution resulted in error.", zap.Error(err))

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    404,
			"message": "Server Error",
		})
	}

	rows, _ = database.DB.Query(context.TODO(), `select iduser from faculty where level='assistant_professor'`)
	var facultyID uuid.UUID
	if rows.Next() {
		rows.Scan(&facultyID)
	}
	insertQuery = "insert into form_to_faculty values($1, $2)"
	res, err = database.DB.Exec(context.TODO(), insertQuery, facultyID, uuidForm)

	if err != nil {
		utils.Logger.Error("Database query execution resulted in error.", zap.Error(err))

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    404,
			"message": "Server Error",
		})
	}

	fmt.Printf("res: %v\n", res)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}
