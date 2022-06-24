package routes

import (
	"admin.server/controllers"
	"admin.server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func userRoutes(app *fiber.App) {
	app.Get("/api/v1/users/:id", controllers.GetUser)

	app.Use(middlewares.AuthenticateMiddleware)

	app.Get("/api/v1/users", controllers.AllUsers)

	app.Post("/api/v1/users", controllers.CreateUser)
}
