package routes

import (
	"github.com/gofiber/fiber/v2"
	"kevinzie/go-commerce/app/controllers"
)

func PublicRoutes(a *fiber.App) {

	route := a.Group("/api/v1")
	route.Post("/login", controllers.SignIn)
}
