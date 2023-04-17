package router

import (
	// swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/handlers/auth"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/handlers/form"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/handlers/home"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/handlers/profile"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
)

func SetupRoutes(app *fiber.App) {
	test := app.Group("/", logger.New())

	// testing route or health check route
	test.Get("", middleware.AuthenticateUser, home.Test)

	// Group api calls with param '/api' : all api routes
	api := app.Group("/api")

	// Group api calls with param '/auth' :autorization routes
	authRoutes := api.Group("/auth")
	authRoutes.Post("/register", auth.Register)
	authRoutes.Post("/signin", auth.SignInUser)
	authRoutes.Get("/refresh", middleware.AuthenticateUser, auth.RefreshAccessToken)
	authRoutes.Post("/signout", middleware.AuthenticateUser, auth.LogoutUser)

	// Group api calls with param '/profile' :user details routes
	profileRoute:=api.Group("/profile", middleware.AuthenticateUser)
	profileRoute.Get("/me", profile.GetMe)
	profileRoute.Get("/dashboard", profile.Dashboard)
	profileRoute.Get("/dashboard/approved", profile.GetApproved)
	profileRoute.Get("/dashboard/rejected", profile.GetRejected)
	profileRoute.Get("/dashboard", profile.Dashboard)

	// Group api calls with param /form$ :form routes
	formRoute:=api.Group("/form", middleware.AuthenticateUser)
	formRoute.Post("/ugintern", form.PostUgIntern)
	formRoute.Get("/ugintern", form.GetUgIntern)
	formRoute.Get("/ugintern/:formID", form.GetForm)
}
