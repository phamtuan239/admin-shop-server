package controllers

import (
	"fmt"
	"strconv"

	"admin.server/database"
	"admin.server/models"
	"github.com/gofiber/fiber/v2"
)

func GetRole(c *fiber.Ctx) error {

	user := models.User{}

	err := database.DB.Preload("Role").First(&user, "id = ?", c.Locals("userId")).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(user)
}

func AllRoles(c *fiber.Ctx) error {
	roles := []models.Role{}

	database.DB.Preload("Permissions").Find(&roles)

	return c.JSON(roles)
}

type RoleDTO struct {
	Name        string
	Permissions []string
}

func CreateRole(c *fiber.Ctx) error {
	roleDto := RoleDTO{}

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	fmt.Println(roleDto)

	list := roleDto.Permissions

	permissions := make([]models.Permission, len(list))

	for i, p := range list {
		pId, _ := strconv.Atoi(p)
		permissions[i] = models.Permission{
			Id: uint(pId),
		}
	}

	role := models.Role{
		Name:        roleDto.Name,
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(true)
}

func UpdateRole(c *fiber.Ctx) error {

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

func DeleteRole(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Delete(&user)

	return c.JSON(true)
}
