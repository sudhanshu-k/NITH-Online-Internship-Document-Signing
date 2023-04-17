package form

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/google/uuid"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func GetUgIntern(c *fiber.Ctx) error {
	user := c.Locals("user").(model.UserResponse)

	getAllForms := "select idform1 from user_to_form where iduser=$1"

	rows, _ := database.DB.Query(context.Background(), getAllForms, user.ID.String())
	utils.FatalError(rows.Err())

	if !rows.Next() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Form not filled."})
	}
	var idform uuid.UUID
	rows.Scan(&idform)
	fmt.Printf("idform: %v\n", idform.String())

	selectQuery := `SELECT name, father_name, address, contact, companyname, areaofintern, "isOffline", "startDay", 
		"endDate", weeks, fromtpo, stipend, formdate, "remakrsDeptt", "remarksFI", id, created_at, updated_at, 
		email from uginternform where id=$1`

	rows, err := database.DB.Query(context.TODO(), selectQuery, idform.String())

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    404,
			"message": "Server Error",
		})
	}

	if !rows.Next() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Form not filled."})
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
