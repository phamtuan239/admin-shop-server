package controllers

import (
	"admin.server/database"
	"admin.server/models"
	"github.com/gofiber/fiber/v2"
)

func AllPerissions(c *fiber.Ctx) error {
	permissions := []models.Permission{}

	database.DB.Find(&permissions)

	return c.Status(200).JSON(permissions)
}
