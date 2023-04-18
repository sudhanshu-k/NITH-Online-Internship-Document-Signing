package profile

import (
	"context"
	"strings"

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
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "No forms filled.", "rows": ""})
		}

		var data uuid.UUID
		rows.Scan(&data)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Filled Form", "rows": data})
	}

	getAllFormsToReview := `select idform1, name, email
			from uginternform, form_to_faculty
			where uginternform.id=form_to_faculty.idform1 and
			form_to_faculty.iduser=$1 and status=$2;`

	rows, err := database.DB.Query(context.Background(), getAllFormsToReview, user.ID.String(), "Pending")
	// utils.FatalError(rows.Err())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "DB Error."})
	}

	var formData []model.FormResponse

	if !rows.Next() {
		return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{"status": "success", "message": "All forms are reviewed.", "rows": ""})
	} else {
		var form model.FormResponse
		rows.Scan(&form.ID, &form.User.Name, &form.User.RollNumber)
		form.FormType = "UG Intern"
		form.User.RollNumber = strings.TrimRight(form.User.RollNumber, "@nith.ac.in")
		formData = append(formData, form)
		for rows.Next() {
			var form model.FormResponse
			rows.Scan(&form.ID, &form.User.Name, &form.User.RollNumber)
			form.FormType = "UG Intern"
			form.User.RollNumber = strings.TrimRight(form.User.RollNumber, "@nith.ac.in")
			formData = append(formData, form)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Forms pending for Review", "rows": formData})
}
