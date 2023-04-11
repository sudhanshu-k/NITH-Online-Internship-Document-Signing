package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/initializers"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	return string(hash)
}

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

func SignInUser(c *fiber.Ctx) error {
	var payload *model.Student

	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	
	fetchUserQuery := `select id, email, password from students where email=$1`
	rows, _ := database.DB.Query(context.Background(), fetchUserQuery, payload.Email)
	utils.FatalError(rows.Err())
	
	if !rows.Next() {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "User doesnot exist."})
	}
	var user model.Student
	rows.Scan(&user.ID, &user.Email, &user.Password)
	
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
	if errAccess != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": errRefresh.Error()})
	}
	fmt.Println(time.Now())
	fmt.Println(time.Now().Add(time.Duration(config.AccessTokenMaxAge)))
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		Secure:   false,
		MaxAge:   config.AccessTokenMaxAge * 60,
		HTTPOnly: true,
		// Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    *refreshTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.RefreshTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		// Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		// Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token})
}

func RefreshAccessToken(c *fiber.Ctx) error {
	message := "could not refresh access token"

	refresh_token := c.Cookies("refresh_token")

	if refresh_token == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	config, _ := initializers.LoadConfig(".")
	ctx := context.TODO()

	tokenClaims, err := utils.ValidateToken(refresh_token, config.RefreshTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	userid, err := initializers.RedisClient.Get(ctx, tokenClaims.TokenUuid).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	var user model.Student

	fetchUserQuery := `select id from students where id=$1`
	rows, _ := database.DB.Query(context.Background(), fetchUserQuery, userid)
	utils.FatalError(rows.Err())
	
	if !rows.Next() {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "User doesnot exist."})
	}
	rows.Scan(&user.ID, &user.Email, &user.Password)

	accessTokenDetails, err := utils.CreateToken(user.ID.String(), time.Duration(config.AccessTokenMaxAge), config.AccessTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	now := time.Now()

	errAccess := initializers.RedisClient.Set(ctx, accessTokenDetails.TokenUuid, user.ID.String(), time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": errAccess.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(config.AccessTokenMaxAge)),
		Secure:   false,
		HTTPOnly: true,
		// Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(config.AccessTokenMaxAge)),
		Secure:   false,
		HTTPOnly: false,
		// Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token})
}

func LogoutUser(c *fiber.Ctx) error {
	message := "Token is invalid or session has expired"

	refresh_token := c.Cookies("refresh_token")

	if refresh_token == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	config, _ := initializers.LoadConfig(".")
	ctx := context.TODO()
	fmt.Println("here")
	// fmt.Println(refresh_token)
	tokenClaims, err := utils.ValidateToken(refresh_token, config.RefreshTokenPublicKey)
	if err != nil {
		fmt.Print("here")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	fmt.Println("here")
	access_token_uuid := c.Locals("access_token_uuid").(string)
	_, err = initializers.RedisClient.Del(ctx, tokenClaims.TokenUuid, access_token_uuid).Result()
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	fmt.Println("here")
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "logged_in",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}