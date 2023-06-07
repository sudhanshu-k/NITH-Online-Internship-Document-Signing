package auth

import (
	"context"
	"time"

	// "github.com/gofiber/redirect/v2"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/config"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func SignInUser(c *fiber.Ctx) error {
	var payload *model.User
	serverError := c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
		"code":    404,
		"message": "Server Error",
	})

	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	fetchUserQuery := `select id, first_name, last_name, email, password, "isFaculty" from users where email=$1`
	rows, _ := database.DB.Query(context.Background(), fetchUserQuery, payload.Email)
	if rows.Err() != nil {
		utils.Logger.Error("Database query execution resulted in error.", zap.Error(rows.Err()))
		return serverError
	}

	if !rows.Next() {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "User doesnot exist."})
	}
	var user model.User
	rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsFaculty)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": "Invalid email or password"})
	}

	accessTokenDetails, err := utils.CreateToken(user.ID.String(), config.Config.AccessTokenExpiresIn, config.Config.AccessTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	refreshTokenDetails, err := utils.CreateToken(user.ID.String(), config.Config.RefreshTokenExpiresIn, config.Config.RefreshTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	ctx := context.TODO()
	now := time.Now()

	err = database.RedisClient.Set(ctx, accessTokenDetails.TokenUuid, user.ID.String(), time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if err != nil {
		utils.Logger.Error("Reddis query execution resulted in error.", zap.Error(err))
		return serverError
	}

	err = database.RedisClient.Set(ctx, refreshTokenDetails.TokenUuid, user.ID.String(), time.Unix(*refreshTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if err != nil {
		utils.Logger.Error("Reddis query execution resulted in error.", zap.Error(err))
		return serverError
	}

	var userData model.UserResponse
	userData.Email = user.Email
	userData.FirstName = user.FirstName
	userData.LastName = user.LastName
	userData.IsLog = true
	userData.IsFaculty = user.IsFaculty

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		Secure:   false,
		MaxAge:   config.Config.AccessTokenMaxAge * 60,
		HTTPOnly: true,
		// Domain:   "localhost",
		// SameSite: "None",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    *refreshTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.Config.RefreshTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		// Domain:   "localhost",
		// SameSite: "None",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.Config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		// Domain:   "localhost",
		// SameSite: "None",
	})

	// return c.Redirect(c.BaseURL() + "/api/profile/me")
	// return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token, "user": userData})
}
