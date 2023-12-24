package errors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadavw1312/golang-fiber/utils"
)

var Errors = &struct {
	InvalidCredentials string
}{
	InvalidCredentials: "Invalid credentials",
}

func GetError(ctx *fiber.Ctx, status int, err string) error {
	return ctx.Status(status).JSON(utils.NewErrorResponse(err))
}
