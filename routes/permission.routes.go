package routes

import (
	"admin.server/controllers"
	"admin.server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func permissionRoutes(app *fiber.App) {
	app.Use(middlewares.AuthenticateMiddleware)

	app.Get("api/v1/permissions", controllers.AllPerissions)

	app.Post("/api/v1/products", controllers.CreateProuct)
}
