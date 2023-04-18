package auth

import (
	"context"
	"time"

	// "github.com/gofiber/redirect/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/initializers"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func LogoutUser(c *fiber.Ctx) error {
	message := "Token is invalid or session has expired"

	refresh_token := c.Cookies("refresh_token")

	if refresh_token == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	config, _ := initializers.LoadConfig(".")
	ctx := context.TODO()
	// fmt.Println("here")
	// fmt.Println(refresh_token)
	tokenClaims, err := utils.ValidateToken(refresh_token, config.RefreshTokenPublicKey)
	if err != nil {
		// fmt.Print("here")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	// fmt.Println("here")
	access_token_uuid := c.Locals("access_token_uuid").(string)
	_, err = initializers.RedisClient.Del(ctx, tokenClaims.TokenUuid, access_token_uuid).Result()
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	// fmt.Println("here")
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Path: "/",
		Expires: expired,
		Secure:   false,
		HTTPOnly: true,
		SameSite: "None",
	})
	c.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Path: "/",
		Expires: expired,
		Secure:   false,
		HTTPOnly: true,
		SameSite: "None",
	})
	c.Cookie(&fiber.Cookie{
		Name:    "logged_in",
		Value:   "",
		Path: "/",
		Expires: expired,
		Secure:   false,
		HTTPOnly: false,
		SameSite: "None",
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
