package auth

import (
	"time"

	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
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
	type CreateUserRequest struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname,omitempty"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	db := database.DB
	json := new(CreateUserRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	password := hashAndSalt([]byte(json.Password))
	err := checkmail.ValidateFormat(json.Email)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Email Address",
		})
	}

	new := model.Student{
		FirstName: json.FirstName,
		LastName:  json.LastName,
		Email:     json.Email,
		Password:  password,
		ID:        uuid.New(),
	}

	found := model.Student{}
	query := model.Student{Email: json.Email}
	err = db.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "User already exists",
		})
	}
	db.Create(&new)
	session := model.Session{UserRefer: new.ID, Sessionid: uuid.New()}
	err = db.Create(&session).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Creation Error",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "sessionid",
		Expires:  time.Now().Add(5 * 24 * time.Hour),
		Value:    session.Sessionid.String(),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    session,
	})
}
