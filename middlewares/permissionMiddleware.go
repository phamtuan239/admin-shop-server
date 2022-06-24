package middlewares

import (
	"admin.server/database"
	"admin.server/models"
	"github.com/gofiber/fiber/v2"
)

func IsPermission(c *fiber.Ctx) {
	id := c.Locals("userId")

	user := models.User{}

	database.DB.Preload("Role").Where("id = ?", id).First(&user)

	
}