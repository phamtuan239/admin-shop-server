package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, page int, entity Entity) fiber.Map {
	limit := 5

	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)

	total := entity.Count(db)

	return fiber.Map{
		"data":  data,
		"total": total,
	}
}
