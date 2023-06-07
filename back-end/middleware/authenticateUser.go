package middleware

import (
	"context"
	// "fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/config"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func AuthenticateUser(c *fiber.Ctx) error {
	var access_token string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		access_token = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("access_token") != "" {
		access_token = c.Cookies("access_token")
	}

	if access_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	tokenClaims, err := utils.ValidateToken(access_token, config.Config.AccessTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	ctx := context.TODO()
	ID, err := database.RedisClient.Get(ctx, tokenClaims.TokenUuid).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "Token is invalid or session has expired"})
	}

	var user model.User
	fetchUserQuery := `select id, first_name, last_name, email, "isFaculty" from users where id=$1`
	rows, _ := database.DB.Query(context.Background(), fetchUserQuery, ID)
	if rows.Err() != nil {
		utils.Logger.Error("Database query execution resulted in error.", zap.Error(rows.Err()))

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    404,
			"message": "Server Error",
		})
	}

	if !rows.Next() {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.IsFaculty)
	userDetails := model.FilterUserRecord(&user)
	userDetails.IsLog = true

	if user.IsFaculty {
		fetchTeacherQuery := `select level from faculty where iduser=$1`
		rows, _ = database.DB.Query(context.Background(), fetchTeacherQuery, userDetails.ID)
		if rows.Err() != nil {
			utils.Logger.Error("Database query execution resulted in error.", zap.Error(rows.Err()))

			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"code":    404,
				"message": "Server Error",
			})
		}

		if !rows.Next() {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Couldn't find entry in corresponding faculty table."})
		} else {
			rows.Scan(&userDetails.Level)
		}
	}

	c.Locals("user", userDetails)
	c.Locals("access_token_uuid", tokenClaims.TokenUuid)

	return c.Next()
}
