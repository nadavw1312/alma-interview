package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadavw1312/golang-fiber/controllers"
)

func InitUserRoutes(api *fiber.App) {
	router := api.Group("/user") // /api
	uc := controllers.UserControllerP

	router.Post("/", uc.CreateUser)

	router.Put("/:id", uc.UpdateById)

	router.Get("/:id", uc.GetUser)
	router.Get("/", uc.GetAll)

	router.Delete("/:id", uc.DeleteUser)
}
