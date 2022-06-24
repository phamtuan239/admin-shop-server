package main

import (
	"log"
	"os"

	"admin.server/database"
	"admin.server/logger"
	"admin.server/routes"
	"github.com/TwiN/go-color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	l := &logger.LoggerCustome{
		InfoLogger:    log.New(os.Stdout, color.InGreen("INFO: "), log.Ldate|log.Ltime|log.Lshortfile),
		WarningLogger: log.New(os.Stdout, color.InYellow("WARNING: "), log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger:   log.New(os.Stdout, color.InRed("ERROR: "), log.Ldate|log.Ltime|log.Lshortfile),
	}

	logger.SetLogger(l)

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000/",
		AllowCredentials: true,
	}))

	routes.SetUp(app)

	app.Listen(":5000")
}
