package main

import (
	"schuett-webapp-api/configs"
	"schuett-webapp-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	app.Listen(":8080")
}
