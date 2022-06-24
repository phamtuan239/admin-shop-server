package controllers

import (
	"strconv"

	"admin.server/database"
	"admin.server/models"
	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {

	userId, _ := strconv.Atoi(c.Params("id"))

	user := models.User{}

	err := database.DB.Preload("Role").First(&user, "id = ?", userId).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(user)
}

func AllProducts(c *fiber.Ctx) error {

	query := QueryPage{}

	c.QueryParser(&query)

	page := query.Page

	return c.JSON(models.Paginate(database.DB, page, &models.Product{}))
}

func CreateProuct(c *fiber.Ctx) error {
	product := models.Product{}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fileHeader, _ := c.FormFile("image")

	file, _ := fileHeader.Open()

	err := product.UploadFile(file)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = database.DB.Create(&product).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(true)
}

func UpdateProduct(c *fiber.Ctx) error {

	userId, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(userId),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(true)
}

func DeleteProduct(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Delete(&user)

	return c.JSON(true)
}
