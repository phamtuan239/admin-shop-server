package controllers

import (
	"strconv"
	"time"

	"admin.server/config"
	"admin.server/database"
	"admin.server/helpers"
	"admin.server/logger"
	"admin.server/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	body := map[string]string{}

	if err := c.BodyParser(&body); err != nil {
		logger.Logger.ErrorLogger.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := models.User{
		FirstName: body["first_name"],
		LastName:  body["last_name"],
		Email:     body["email"],
		Password:  body["password"],
		RoleId:    1,
	}

	user.SetPassword()

	err := database.DB.Create(&user).Error

	if err != nil {
		logger.Logger.ErrorLogger.Println(err)
		return c.Status(403).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(true)
}

func Login(c *fiber.Ctx) error {
	body := map[string]string{}

	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	user := models.User{}

	err = database.DB.First(&user, "email = ?", body["email"]).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Email not found",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body["password"]))

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Password not match",
		})
	}

	issure := strconv.Itoa(int(user.Id))

	timeAccess := time.Now().Add(config.AccessTokenExpire).Unix()

	timeRefresh := time.Now().Add(config.RefreshCookieExpire).Unix()

	accessToken, err := helpers.GenerateJWT(issure, timeAccess, config.Secret())

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	refreshToken, err := helpers.GenerateJWT(issure, timeRefresh, config.Secret())

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	accessCookie := fiber.Cookie{
		Name:     "access",
		Value:    accessToken,
		Expires:  time.Now().Add(config.AccessCookieExpire),
		HTTPOnly: true,
	}

	refreshCookie := fiber.Cookie{
		Name:     "refresh",
		Value:    refreshToken,
		Expires:  time.Now().Add(config.RefreshCookieExpire),
		HTTPOnly: true,
	}

	c.Cookie(&accessCookie)

	c.Cookie(&refreshCookie)

	return c.JSON(fiber.Map{
		"message": "sucess",
	})
}

func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(time.Minute),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func GetInfo(c *fiber.Ctx) error {
	id := c.Locals("userId")

	user := models.User{}

	database.DB.Preload("Role").Preload("Permissions").Where("id = ?", id).First(&user)

	return c.Status(200).JSON(fiber.Map{
		"user": user,
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	body := map[string]string{}

	id := c.Locals("userId")

	c.BodyParser(&body)

	user := models.User{}

	database.DB.Where("id = ?", id).First(&user).Updates(body)

	return c.JSON(true)
}

func UpdatePassword(c *fiber.Ctx) error {
	body := map[string]string{}
	id := c.Locals("userId")

	c.BodyParser(&body)
	user := models.User{}
	
	database.DB.Where("id = ?", id).First(&user).Updates(body)
	return c.JSON(true)
}
