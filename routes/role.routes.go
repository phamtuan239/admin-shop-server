package routes

import (
	"admin.server/controllers"
	"admin.server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func roleRoutes(app *fiber.App) {

	app.Use(middlewares.AuthenticateMiddleware)
	
	app.Post("/api/v1/roles", controllers.CreateRole)

	app.Get("/api/v1/roles", controllers.AllRoles)
}
