package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kevinzie/go-commerce/app/service"
)

type AuthController interface {
	SignUp(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
}

func SignIn(c *fiber.Ctx) error {
	return c.Status(200).JSON(service.Login(c))
}
