package middleware

import (
	"context"
	// "fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/initializers"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func AuthenticateUser(c *fiber.Ctx) error {
	var access_token string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		// fmt.Print(access_token + "sdf")
		access_token = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("access_token") != "" {
		access_token = c.Cookies("access_token")
	}

	if access_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	config, _ := initializers.LoadConfig(".")

	tokenClaims, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	ctx := context.TODO()
	ID, err := initializers.RedisClient.Get(ctx, tokenClaims.TokenUuid).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "Token is invalid or session has expired"})
	}
	// fmt.Print(ID)

	var user model.User
	fetchUserQuery := `select id, first_name, last_name, email, "isFaculty" from users where id=$1`
	rows, _ := database.DB.Query(context.Background(), fetchUserQuery, ID)
	utils.FatalError(rows.Err())
	// fmt.Println("here")

	if !rows.Next() {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.IsFaculty)
	userDetails := model.FilterUserRecord(&user)

	if user.IsFaculty {
		fetchTeacherQuery := `select level from faculty where iduser=$1`
		rows, _ = database.DB.Query(context.Background(), fetchTeacherQuery, userDetails.ID)
		if !rows.Next() {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Couldn't find entry in corresponding faculty table."})
		} else {
			rows.Scan(&userDetails.Level)
		}
	}

	c.Locals("user", userDetails)
	c.Locals("access_token_uuid", tokenClaims.TokenUuid)

	// fmt.Printf("here")
	return c.Next()
}
