package routes

import (
	"github.com/gofiber/fiber/v2"
	"kevinzie/go-commerce/app/controllers"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.

	//authz := fibercasbin.New(fibercasbin.Config{
	//	ModelFilePath: "./auth_model.conf",
	//	//PolicyAdapter: xormadapter.NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/"),
	//	Lookup: func(c *fiber.Ctx) string {
	//		// fetch authenticated user subject
	//		log.Printf("test", c)
	//		return ""
	//	},
	//})

	route := a.Group("/api/v1")

	route.Get("/users", controllers.GetUsers)
	route.Get("/user/:id", controllers.GetUser)
	route.Post("/user", controllers.CreateUser)
	route.Put("/user/:id", controllers.UpdateUser)
	route.Delete("/user/:id", controllers.DeleteUser)
}
