package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadavw1312/golang-fiber/controllers"
	"github.com/nadavw1312/golang-fiber/middleware"
)

func InitUserRoutes(api *fiber.App) {
	user := api.Group("/user") // /api
	uc := controllers.NewUserController("users")

	user.Post("/", uc.CreateUser)
	user.Post("/signin", uc.Signin)

	user.Get("/protected", middleware.Protected(), func(c *fiber.Ctx) error {
		return c.SendString("You are protected")
	})
	user.Get("/:id", uc.GetUser)

	user.Delete("/:id", uc.DeleteUser)

}
