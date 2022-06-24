package routes

import (
	"admin.server/controllers"
	"admin.server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func authRoutes(app *fiber.App) {
	app.Post("/api/v1/auth/register", controllers.Register)

	app.Post("/api/v1/auth/login", controllers.Login)

	app.Use(middlewares.AuthenticateMiddleware)

	app.Get("/api/v1/auth/info", controllers.GetInfo)

	app.Get("/api/v1/auth/logout", controllers.Logout)
}
