package profile

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func GetApproved(c *fiber.Ctx) error {
	user := c.Locals("user").(model.UserResponse)

	if !user.IsFaculty {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "failed", "message": "Not faculty."})
	}

	getAllForms := `select idform1, name, email
			from uginternform, form_to_faculty
			where uginternform.id=form_to_faculty.idform1 and
			form_to_faculty.iduser=$1 and status=$2;`

	rows, err := database.DB.Query(context.Background(), getAllForms, user.ID.String(), "Approved")
	// utils.FatalError(rows.Err())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "DB Error."})
	}

	var formData []model.FormResponse

	if !rows.Next() {
		return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{"status": "success", "message": "No approved forms.", "rows": ""})
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Approved Forms", "rows": formData})
}
