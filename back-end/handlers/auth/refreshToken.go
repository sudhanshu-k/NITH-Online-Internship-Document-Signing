package auth

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/config"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func RefreshAccessToken(c *fiber.Ctx) error {
	message := "could not refresh access token"
	serverError := c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
		"code":    404,
		"message": "Server Error",
	})

	refresh_token := c.Cookies("refresh_token")
	if refresh_token == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	ctx := context.TODO()

	tokenClaims, err := utils.ValidateToken(refresh_token, config.Config.RefreshTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	userid, err := database.RedisClient.Get(ctx, tokenClaims.TokenUuid).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	} else if err != nil {
		utils.Logger.Error("Reddis query execution resulted in error.", zap.Error(err))
		return serverError
	}

	var user model.User

	fetchUserQuery := `select id from users where id=$1`
	rows, _ := database.DB.Query(context.Background(), fetchUserQuery, userid)
	if rows.Err() != nil {
		utils.Logger.Error("Database query execution resulted in error.", zap.Error(rows.Err()))
		return serverError
	}

	if !rows.Next() {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "User doesnot exist."})
	}
	rows.Scan(&user.ID)

	accessTokenDetails, err := utils.CreateToken(user.ID.String(), config.Config.AccessTokenExpiresIn, config.Config.AccessTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	now := time.Now()

	errAccess := database.RedisClient.Set(ctx, accessTokenDetails.TokenUuid, user.ID.String(), time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": errAccess.Error()})
	}

	// println(*accessTokenDetails.Token)
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
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.Config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		// Domain:   "localhost",
		// SameSite: "None",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token})
}
