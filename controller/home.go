package controller

import (
	"belajar/entity"
	"belajar/model"
	"belajar/repository"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {

	return c.SendString("Hello, World!")
}

func GetUser(c *fiber.Ctx) error {

	user, err := repository.GetUser()
	if err != nil {
		return c.JSON(model.ResponseMessage{Message: "Data tidak ditemukan"})
	}

	return c.JSON(model.Response[entity.User]{user})
}
