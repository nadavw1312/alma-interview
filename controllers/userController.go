package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nadavw1312/golang-fiber/dal"
	"github.com/nadavw1312/golang-fiber/database"
	"github.com/nadavw1312/golang-fiber/env"
	"github.com/nadavw1312/golang-fiber/models"
	"github.com/nadavw1312/golang-fiber/types"
	"github.com/nadavw1312/golang-fiber/utils"
	"github.com/nadavw1312/golang-fiber/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userDal *dal.UserDal
}

func NewUserController(collectionName string) *UserController {
	userDal := dal.NewUserDal(database.MG.Db.Collection(collectionName))

	return &UserController{userDal}
}

func (uc *UserController) Signin(ctx *fiber.Ctx) error {
	var credentials types.Credentials
	if err := ctx.BodyParser(&credentials); err != nil {
		return errors.GetError(ctx, http.StatusBadRequest, err.Error())
	}

	user, err := uc.userDal.FindByUsername(ctx.Context(), credentials.Username)
	if err != nil {
		return errors.GetError(ctx, http.StatusBadRequest, errors.Errors.InvalidCredentials)
	}

	if err := comparePassword(credentials.Password, user.Password); err != nil {
		return errors.GetError(ctx, http.StatusBadRequest, errors.Errors.InvalidCredentials)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.Id
	claims["userName"] = user.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	t, err := token.SignedString([]byte(env.Env.AppSecret))
	if err != nil {
		return errors.GetError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

func (uc *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := uc.userDal.GetById(ctx.Context(), id)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	return ctx.JSON(user)
}

func (uc *UserController) CreateUser(ctx *fiber.Ctx) error {
	newUser := models.CreateUserRequest{}

	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}

	// Validation
	if errs := utils.Validator.Validate(newUser); len(errs.Errors) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(errs)
	}

	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err.Error()))
	}
	newUser.Password = hashedPassword

	userId, err := uc.userDal.InsertUser(ctx.Context(), newUser)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
	}

	createdUser, err := uc.userDal.GetById(ctx.Context(), userId)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err.Error()))
	}

	return ctx.Status(201).JSON(createdUser)

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
