package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nadavw1312/golang-fiber/dal"
	"github.com/nadavw1312/golang-fiber/models"
	"github.com/nadavw1312/golang-fiber/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userDal *dal.UserDal
}

func NewUserController() *UserController {
	userDal := dal.NewUserDal()

	return &UserController{userDal}
}

func (userController *UserController) GetAll(ctx *fiber.Ctx) error {
	users, err := userController.userDal.GetAll()
	if err != nil {
		return err
	}
	return ctx.JSON(users)
}

func (uc *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := uc.userDal.GetById(id)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	return ctx.JSON(user)
}

func (uc *UserController) GetById(id string) (models.User, error) {
	return uc.userDal.GetById(id)
}

func (uc *UserController) CreateUser(ctx *fiber.Ctx) error {
	newUser := models.CreateUserRequest{}

	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	userId, err := uc.userDal.InsertUser(ctx.Context(), newUser)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
	}

	createdUser, err := uc.userDal.GetById(userId)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
	}

	return ctx.Status(201).JSON(createdUser)
}

func (uc *UserController) UpdateById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	newUser := models.User{}

	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}
	newUser.Id = id

	if err := uc.userDal.UpdateById(ctx.Context(), id, newUser); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
	}

	return ctx.SendString(fmt.Sprintf("Updated user %s", id))
}

func (uc *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := uc.userDal.DeleteById(ctx.Context(), id); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	return ctx.SendString(fmt.Sprintf("Deleted user %s", id))
}

func hashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password with a cost of 14 (you can adjust the cost as needed)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
