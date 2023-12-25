package controllers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nadavw1312/golang-fiber/dal"
	"github.com/nadavw1312/golang-fiber/models"
	"github.com/nadavw1312/golang-fiber/utils"
)

type FollowUsersController struct {
	dal *dal.FollowUserDal
}

func NewFollowUsersController() *FollowUsersController {
	followUserDal := dal.NewFollowUserDal()

	return &FollowUsersController{followUserDal}
}

func (controller *FollowUsersController) Follow(ctx *fiber.Ctx) error {
	newFollowUser := models.NewFollowUser{}

	if err := ctx.BodyParser(&newFollowUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	if err := validateFollowersIds(newFollowUser.FollowerId, newFollowUser.FollowedId); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	if err := controller.validateFollowerNotExists(newFollowUser.FollowerId, newFollowUser.FollowedId); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	id, err := controller.dal.Follow(newFollowUser.FollowerId, newFollowUser.FollowedId)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
	}
	createdUser, err := controller.dal.GetById(id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
	}

	return ctx.Status(201).JSON(createdUser)

}

func (controller *FollowUsersController) Unfollow(ctx *fiber.Ctx) error {
	body := models.NewFollowUser{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	if err := validateFollowersIds(body.FollowerId, body.FollowedId); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	return controller.dal.Unfollow(body.FollowerId, body.FollowedId)
}

func (controller *FollowUsersController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	followUser, err := controller.dal.GetById(id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
	}

	return ctx.Status(200).JSON(followUser)
}

func (controller *FollowUsersController) GetUserFollowersIds(followerId string) []string {
	followersIds := []string{}
	userFollowers, err := controller.dal.GetByFollowerId(followerId)
	if err != nil {
		return nil
	}

	for _, follower := range userFollowers {
		followersIds = append(followersIds, follower.FollowedId)
	}

	return utils.RemoveDuplicates(followersIds)
}

func validateFollowersIds(followerId string, followedId string) error {
	if _, err := UserControllerP.GetById(followedId); err != nil {
		return errors.New("Followed user id does not exist")
	}

	if _, err := UserControllerP.GetById(followerId); err != nil {
		return errors.New("Follower user id does not exist")
	}

	return nil
}

func (controller *FollowUsersController) validateFollowerNotExists(followerId string, followId string) error {
	followUsers, err := controller.dal.GetByFollowerId(followerId)
	if err != nil {
		return err
	}

	for _, followUser := range followUsers {
		if followUser.FollowedId == followId {
			return errors.New("Follower already exists")
		}
	}

	return nil
}
