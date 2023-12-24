package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nadavw1312/golang-fiber/database"
	"github.com/nadavw1312/golang-fiber/env"
	"github.com/nadavw1312/golang-fiber/routes"
)

func main() {
	// Default config
	app := fiber.New()
	env.LoadEnv()
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:3000",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	database.Connect()
	// routes
	routes.InitUserRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
