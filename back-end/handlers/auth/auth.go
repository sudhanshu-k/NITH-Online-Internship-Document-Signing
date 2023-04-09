package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	return string(hash)
}

// func comparePasswords(hashedPwd string, plainPwd []byte) bool {
// 	byteHash := []byte(hashedPwd)
// 	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
// 	return err == nil
// }

// Universal date the Session Will Expire
func SessionExpires() time.Time {
	return time.Now().Add(5 * 24 * time.Hour)
}

func Register(c *fiber.Ctx) error {
	db := database.DB
	json := new(model.Student)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	json.Password = hashAndSalt([]byte(json.Password))
	err := checkmail.ValidateFormat(json.Email)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Email Address",
		})
	}

	insertQuery := `insert into students values($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	res, err := db.Exec(context.Background(), insertQuery, uuid.New(), json.FirstName, json.LastName, json.Email, json.Password)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "User already exists",
		})
	}
	fmt.Printf("res: %v\n", res)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
	
}
