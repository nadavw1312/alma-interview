package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nadavw1312/golang-fiber/controllers"
	"github.com/nadavw1312/golang-fiber/routes"
)

func main() {
	// Default config
	app := fiber.New()
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	controllers.InitControllers()

	// routes
	routes.InitUserRoutes(app)
	routes.InitFollowUserRoutes(app)
	routes.InitTweetRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
