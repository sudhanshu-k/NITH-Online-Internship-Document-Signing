package form

import (
	"context"
	// "fmt"
	// "time"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/proxy"
	// "github.com/google/uuid"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func GetForm(c *fiber.Ctx) error {
	user := c.Locals("user").(model.UserResponse)

	if !user.IsFaculty{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "failed", "message": "Faculty only route."})
	}
	selectQuery := `SELECT name, father_name, address, contact, companyname, areaofintern, "isOffline", "startDay", 
		"endDate", weeks, fromtpo, stipend, formdate, "remakrsDeptt", "remarksFI", id, created_at, updated_at, 
		email from uginternform where id=$1`

	rows, err := database.DB.Query(context.TODO(), selectQuery, c.Params("formID"))

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    404,
			"message": "Server Error",
		})
	}

	if !rows.Next() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Wrong form id."})
	}

	var data model.UGInternForm
	rows.Scan(&data.Name, &data.FatherName, &data.Address, &data.Contact,
		&data.CompanyName, &data.AreaOfIntrest, &data.IsOffline, &data.StartDay, &data.EndDay,
		&data.Weeks, &data.FromTPO, &data.Stipend, &data.FormDate, &data.RemarksDept, &data.RemarksFI,
		&data.ID, &data.CreatedAt, &data.UpdatedAt, &data.Email)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"formData": data,
	})
}
