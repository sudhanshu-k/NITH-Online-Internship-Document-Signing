package main

import (
	// errors "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/router"

	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//new fiber instance
	app := fiber.New()

	//connect to db
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
