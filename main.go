package main

import (
	"belajar/config"
	"belajar/controller"

	"github.com/gofiber/fiber/v2"
)

func InitClean() {

	app := fiber.New()

	app.Get("/", controller.Home)
	app.Get("/user", controller.GetUser)

	app.Listen(":9000")
}

func Init() {

	app := fiber.New()

	app.Get("/", controller.Home)
	app.Get("/user", controller.GetUser)

	app.Listen(":9000")
}

func main() {

	config.ConnectDB()
	Init()
}
