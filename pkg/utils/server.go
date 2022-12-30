package utils

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	fiberConnections, _ := ConnectionURLBuilder("fiber")
	log.Printf("FIBER CONNECTION", fiberConnections)
	if err := a.Listen(fiberConnections); err != nil {
		log.Printf("Ooops... Server is not running!!! Reason: %v", err)
	}
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnection := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnection)
	}()

	fiberConnectionURL, _ := ConnectionURLBuilder("fiber")

	// Running Server
	if err := a.Listen(fiberConnectionURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
	log.Printf(fiberConnectionURL)
	<-idleConnection
}
