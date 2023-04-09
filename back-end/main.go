package main

import (
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/initializers"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/router"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	config, err := initializers.LoadConfig(".")
	middleware.LogIfError(err, "Failed to load environment variables! \n")

	initializers.ConnectDB(&config)
	initializers.ConnectRedis(&config)
}

func main() {
	//new fiber instance
	app := fiber.New()

	// Setup the router
	router.SetupRoutes(app)

	config, err := initializers.LoadConfig(".")
	middleware.LogIfError(err, "Failed to load environment variables! \n")
	
	app.Listen(config.PORT)
}
