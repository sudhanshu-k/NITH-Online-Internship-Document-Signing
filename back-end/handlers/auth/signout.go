package auth

import (
	"context"
	"time"

	// "github.com/gofiber/redirect/v2"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/config"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func LogoutUser(c *fiber.Ctx) error {
	message := "Token is invalid or session has expired"

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
		// fmt.Print("here")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	access_token_uuid := c.Locals("access_token_uuid").(string)
	_, err = database.RedisClient.Del(ctx, tokenClaims.TokenUuid, access_token_uuid).Result()
	if err != nil {
		utils.Logger.Error("Reddis query execution resulted in error.", zap.Error(err))
		return serverError
	}

	// fmt.Println("here")
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		Expires:  expired,
		Secure:   false,
		HTTPOnly: true,
		// SameSite: "None",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		Expires:  expired,
		Secure:   false,
		HTTPOnly: true,
		// SameSite: "None",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "",
		Path:     "/",
		Expires:  expired,
		Secure:   false,
		HTTPOnly: false,
		// SameSite: "None",
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
