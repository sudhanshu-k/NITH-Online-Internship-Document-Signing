package auth

import (
	"context"
	"time"

	// "github.com/gofiber/redirect/v2"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/initializers"
	"golang.org/x/crypto/bcrypt"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func SignInUser(c *fiber.Ctx) error {
	var payload *model.User

	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	fetchUserQuery := `select id, first_name, last_name, email, password, "isFaculty" from users where email=$1`
	rows, _ := database.DB.Query(context.Background(), fetchUserQuery, payload.Email)
	utils.FatalError(rows.Err())

	if !rows.Next() {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "User doesnot exist."})
	}
	var user model.User
	rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsFaculty)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": "Invalid email or password"})
	}

	config, _ := initializers.LoadConfig(".")
	// fmt.Println(config.AccessTokenPrivateKey)
	accessTokenDetails, err := utils.CreateToken(user.ID.String(), config.AccessTokenExpiresIn, config.AccessTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	refreshTokenDetails, err := utils.CreateToken(user.ID.String(), config.RefreshTokenExpiresIn, config.RefreshTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	ctx := context.TODO()
	now := time.Now()

	errAccess := initializers.RedisClient.Set(ctx, accessTokenDetails.TokenUuid, user.ID.String(), time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": errAccess.Error()})
	}

	errRefresh := initializers.RedisClient.Set(ctx, refreshTokenDetails.TokenUuid, user.ID.String(), time.Unix(*refreshTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errRefresh != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": errRefresh.Error()})
	}
	// fmt.Println(time.Now())
	// fmt.Println(time.Now().Add(time.Duration(config.AccessTokenMaxAge)))

	var userData model.UserResponse
	userData.Email = user.Email
	userData.FirstName = user.FirstName
	userData.LastName = user.LastName
	userData.IsLog = true
	userData.IsFaculty=user.IsFaculty

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		Secure:   false,
		MaxAge:   config.AccessTokenMaxAge * 60,
		HTTPOnly: true,
		// Domain:   "localhost",
		SameSite: "None",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    *refreshTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.RefreshTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		// Domain:   "localhost",
		SameSite: "None",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		// Domain:   "localhost",
		SameSite: "None",
	})

	// return c.Redirect(c.BaseURL() + "/api/profile/me")
	// return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token, "user": userData})
}
