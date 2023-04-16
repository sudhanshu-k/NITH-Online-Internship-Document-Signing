package form

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/google/uuid"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func FillUgIntern(c *fiber.Ctx) error {
	user := c.Locals("user").(model.UserResponse)

	getAllForms := "select idform1 from user_to_form where iduser=$1"

	rows, _ := database.DB.Query(context.Background(), getAllForms, user.ID.String())
	utils.FatalError(rows.Err())

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
		$15, $16, $17, $18, $19, $20)`
	res, err := database.DB.Exec(context.TODO(), insertQuery, formData.Name, formData.FatherName, formData.Address, formData.Contact,
		formData.CompanyName, formData.AreaOfIntrest, formData.IsOffline, formData.StartDay, formData.EndDay,
		formData.Weeks, formData.FromTPO, formData.Stipend, formData.FormDate, formData.RemarksDept, formData.RemarksFI, 0,
		uuidForm, time.Now(), time.Now(), formData.Email)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    404,
			"message": "Server Error",
		})
	}

	fmt.Printf("res: %v\n", res)
	fmt.Printf("user.ID: %v\n", user.ID.String())
	fmt.Printf("user.ID: %v\n", uuidForm.String())
	insertQuery = "insert into user_to_form values($1, $2)"
	res, err = database.DB.Exec(context.TODO(), insertQuery, user.ID.String(), uuidForm.String())

	

	if err != nil {
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
