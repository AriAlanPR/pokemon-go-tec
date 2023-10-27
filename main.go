package main

import (
	"gofiber/config"
	"gofiber/database"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	database.Init()
	config.SetupRoutes(app)
	app.Listen(3000)

	defer database.Close()
}
