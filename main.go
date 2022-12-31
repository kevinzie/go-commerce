package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // Load env automatically
	"kevinzie/go-commerce/pkg/config"
	"kevinzie/go-commerce/pkg/routes"
	"kevinzie/go-commerce/pkg/utils"
	"os"
)

func main() {
	fiberConfig := config.FiberConfig()
	app := fiber.New(fiberConfig)
	//database.Config

	err := config.Connect()
	if err != nil {
		return
	}
	//middleware.Fiber
	routes.SwaggerRoute(app)
	routes.PrivateRoutes(app)
	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	if os.Getenv("APP_ENV") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
