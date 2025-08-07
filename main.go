package main

import (
	"belajar/config"
	"belajar/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.ConnectDB()

	app := fiber.New()

	router.RouteHome(app)

	app.Listen(":9000")
}
