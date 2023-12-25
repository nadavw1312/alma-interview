package controllers

import (
	"net/http"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/nadavw1312/golang-fiber/dal"
	"github.com/nadavw1312/golang-fiber/models"
	"github.com/nadavw1312/golang-fiber/utils"
)

type TweetController struct {
	dal *dal.TweetDal
}

func NewTweetController() *TweetController {
	tweetDal := dal.NewTweetDal()

	return &TweetController{tweetDal}
}

func (controller *TweetController) InsertTweet(ctx *fiber.Ctx) (F error) {
	tweet := models.NewTweet{}
	if err := ctx.BodyParser(&tweet); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	if _, err := UserControllerP.GetById(tweet.UserId); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse("Followed user id does not exist"))

	}

	tweetId, err := controller.dal.InsertTweet(tweet)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
	}

	createdUser, err := controller.dal.GetTweetById(tweetId)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
	}

	return ctx.Status(201).JSON(createdUser)
}

func (controller *TweetController) GetMostRecentTweetsUserFollows(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	if _, err := UserControllerP.GetById(userId); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse("user id does not exist"))

	}

	followerUsersIds := FollowUsersControllerP.GetUserFollowersIds(userId)
	followsTweets := []*models.Tweet{}

	for _, followerUserId := range followerUsersIds {
		userTweets, err := controller.dal.GetTweetsByUserId(followerUserId)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
		}

		followsTweets = append(followsTweets, userTweets...)
	}

	// Sort the slice by the Date property
	sort.Sort(models.TweetsByDate(followsTweets))

	return ctx.Status(200).JSON(followsTweets)
}
