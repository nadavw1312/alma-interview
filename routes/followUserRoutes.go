package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadavw1312/golang-fiber/controllers"
)

func InitFollowUserRoutes(api *fiber.App) {
	router := api.Group("/followUsers") // /api
	controller := controllers.FollowUsersControllerP

	router.Post("/follow", controller.Follow)
	router.Post("/unFollow", controller.Unfollow)

	router.Get("/:id", controller.GetById)

}
