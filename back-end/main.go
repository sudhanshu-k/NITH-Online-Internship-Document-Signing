package main

import (
	// "log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	// "gorm.io/gorm"
)

// func handleError(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error{
		err:=c.SendString("Api Is Up, Testing Route");
		return err
	})

	app.Listen(":8000")
}
