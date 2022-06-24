package controllers

import (
	"strconv"

	"admin.server/database"
	"admin.server/models"
	"github.com/dgrijalva/jwt-Go"
	"github.com/gofiber/fiber/v2"
)

type Claims struct {
	jwt.StandardClaims
}

type QueryPage struct {
	Page int
}

func GetUser(c *fiber.Ctx) error {

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

func AllUsers(c *fiber.Ctx) error {


	query := QueryPage{}

	c.QueryParser(&query)

	page := query.Page

	return c.JSON(models.Paginate(database.DB, page, &models.User{}))
}

func CreateUser(c *fiber.Ctx) error {
	body := map[string]string{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	role_id, _ := strconv.Atoi(body["role_id"])

	user := models.User{
		FirstName: body["first_name"],
		LastName:  body["last_name"],
		Email:     body["email"],
		Password:  body["password"],
		RoleId:    uint(role_id),
	}

	user.SetPassword()

	database.DB.Create(&user)

	return c.JSON(true)
}

func UpdateUser(c *fiber.Ctx) error {

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

func DeleteUser(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Delete(&user)

	return c.JSON(true)
}
