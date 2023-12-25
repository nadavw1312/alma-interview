package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadavw1312/golang-fiber/controllers"
)

func InitTweetRoutes(api *fiber.App) {
	router := api.Group("/tweet") // /api
	controller := controllers.TweetControllerP

	router.Post("/", controller.InsertTweet)
	router.Get("/getMostRecentTweetsByUserId/:userId", controller.GetMostRecentTweetsUserFollows)
}
