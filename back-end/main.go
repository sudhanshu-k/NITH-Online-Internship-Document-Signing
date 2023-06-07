package main

import (
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/initializers"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/router"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/config"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func init() {
	utils.Logger = zap.Must(zap.NewProduction())

	initializers.LoadConfig()
	initializers.ConnectDB()
	initializers.ConnectRedis()
}

func main() {
	//new fiber instance
	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowHeaders:     "Origin, Content-Type, Accept",
	// 	AllowCredentials: true,
	// }))

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://localhost:5173",
	// 	AllowCredentials: true,
	// }))

	// app.Use(middleware.Security)

	// Setup the router
	router.SetupRoutes(app)

	app.Listen(config.Config.PORT)
}
