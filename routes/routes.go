package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	authRoutes(app)

	userRoutes(app)

	roleRoutes(app)

	permissionRoutes(app)
}
