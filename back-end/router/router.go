package router

import (
	// swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/handlers/auth"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/handlers/home"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
)

func SetupRoutes(app *fiber.App) {
	// app.Get("/swagger/*", swagger.Handler)

	// Setup test routes, can use same syntax to add routes for more models
	// for testing app
	test := app.Group("/", logger.New())
	test.Get("", middleware.AuthenticateUser ,home.Test)

	// Group api calls with param '/api'
	api := app.Group("/api")

	authRoutes := api.Group("/auth")
	authRoutes.Post("/register", auth.Register)
	authRoutes.Post("/signin", auth.SignInUser)
	authRoutes.Post("/refresh", middleware.AuthenticateUser, auth.RefreshAccessToken)
	authRoutes.Post("/signout", middleware.AuthenticateUser, auth.LogoutUser)

	authRoutes.Get("/users/me", middleware.AuthenticateUser, auth.GetMe)
}
