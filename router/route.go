package router

import (
	"belajar/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteHome(app *fiber.App) {

	app.Get("/", controller.Home)
	app.Get("/user", controller.GetUser)
}
