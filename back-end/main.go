package main

import (
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/initializers"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/router"

	"github.com/gofiber/fiber/v2"
)

func init() {
	config, err := initializers.LoadConfig(".")
	utils.LogIfError(err, "Failed to load environment variables! \n")

	initializers.ConnectDB(&config)
	initializers.ConnectRedis(&config)
}

func main() {
	//new fiber instance
	app := fiber.New()

	// Setup the router
	router.SetupRoutes(app)

	config, err := initializers.LoadConfig(".")
	utils.LogIfError(err, "Failed to load environment variables! \n")

	app.Listen(config.PORT)
}
