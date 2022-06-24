package middlewares

import (
	"time"

	"admin.server/config"
	"admin.server/helpers"
	"admin.server/logger"
	"github.com/gofiber/fiber/v2"
)

func AuthenticateMiddleware(c *fiber.Ctx) error {
	accessCookie := c.Cookies("access")

	refreshCookie := c.Cookies("refresh")

	if len(accessCookie) != 0 {
		claims, err := helpers.ParseJWT(accessCookie, config.Secret())

		if err != nil {
			logger.Logger.ErrorLogger.Println(err)
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthenticated",
			})
		}

		c.Locals("userId", claims.Issuer)

		return c.Next()
	}

	if len(refreshCookie) != 0 {
		claims, err := helpers.ParseJWT(refreshCookie, config.Secret())

		if err != nil {
			logger.Logger.ErrorLogger.Println(err)
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthenticated",
			})
		}

		timeAccess := time.Now().Add(config.AccessTokenExpire).Unix()

		accessToken, err := helpers.GenerateJWT(claims.Issuer, timeAccess, config.Secret())

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		accessCookie := fiber.Cookie{
			Name:     "access",
			Value:    accessToken,
			Expires:  time.Now().Add(config.AccessCookieExpire),
			HTTPOnly: true,
		}

		c.Cookie(&accessCookie)

		c.Locals("userId", claims.Issuer)

		return c.Next()
	}

	return c.Status(401).JSON(fiber.Map{
		"message": "Cookie expired",
	})
}
